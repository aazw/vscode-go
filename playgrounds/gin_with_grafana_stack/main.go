package main

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v3"

	// Profiling
	"github.com/grafana/pyroscope-go"

	// OpenMetrics (Prometheus)
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// Tracing
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

const (
	appName         = "ginapp"
	appEnvVarPrefix = "GINAPP_"
)

var (
	handlerOptions = &slog.HandlerOptions{
		AddSource: true, // 行番号などを付与
	}
	sharedLoggerOut = os.Stderr
	sharedLogger    *slog.Logger
)

var tracer = otel.Tracer("ginapp")

var (
	loggerFormat  string
	host          string
	port          int
	pyroscopeURL  string // Grafana Pyroscope
	tempoHostPort string // Grafana Tempo

	cmd = &cli.Command{
		Name:  appName,
		Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "logger",
				Sources:     cli.EnvVars(appEnvVarPrefix + "LOGGER"),
				Usage:       "logger = (text|json)",
				Value:       "text",
				Destination: &loggerFormat,
				Action: func(ctx context.Context, cmd *cli.Command, s string) error {
					switch strings.ToLower(s) {
					case "text", "json":
						return nil
					default:
						return fmt.Errorf("invalid argument: logger is either\"text\" or \"json\", but the value is \"%s\"", s)
					}
				},
			},
			&cli.StringFlag{
				Name:        "host",
				Sources:     cli.EnvVars(appEnvVarPrefix + "HOST"),
				Usage:       "",
				Value:       "localhost",
				Destination: &host,
				Action: func(ctx context.Context, cmd *cli.Command, s string) error {
					if s == "localhost" {
						return nil
					}
					if ip := net.ParseIP(s); ip == nil {
						return nil
					}
					return fmt.Errorf("invalid argument: host is either \"localhost\" or IPv4 or IPv6, but the value is \"%s\"", s)
				},
			},
			&cli.IntFlag{
				Name:        "port",
				Sources:     cli.EnvVars(appEnvVarPrefix + "PORT"),
				Usage:       "",
				Value:       8080,
				Destination: &port,
				Action: func(ctx context.Context, cmd *cli.Command, v int) error {
					if v < 0 || v >= 65536 {
						return fmt.Errorf("invalid argument: port is between 0 and 65535, but the value is %d", v)
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "pyroscope_url",
				Sources: cli.EnvVars(appEnvVarPrefix + "PYROSCOPE_URL"),
				Usage:   "",
				Value:   "http://localhost:4040",
				// Value: "http://pyroscope:4040",
				// Value: "http://host.docker.internal:4040",
				Destination: &pyroscopeURL,
				Action: func(ctx context.Context, cmd *cli.Command, s string) error {
					_, err := url.Parse(s)
					if err != nil {
						return fmt.Errorf("invalid argument: pyroscope_url is \"<scheme>://<host>:<port>\", the value is \"%s\"", s)
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:    "tempo_url",
				Sources: cli.EnvVars(appEnvVarPrefix + "TEMPO_URL"),
				Usage:   "",
				Value:   "localhost:4318",
				// Value: "tempo:4318",
				// Value: "host.docker.internal:4318",
				Destination: &tempoHostPort,
				Action: func(ctx context.Context, cmd *cli.Command, s string) error {
					_, _, err := net.SplitHostPort(s)
					if err != nil {
						return fmt.Errorf("invalid argument: tempo_url is \"<host>:<port>\", the value is \"%s\"", s)
					}
					return nil
				},
			},
		},
		Action: run,
	}
)

func init() {
	handler := slog.NewTextHandler(sharedLoggerOut, handlerOptions)
	sharedLogger = slog.New(handler)
}

func main() {
	args := os.Args[1:]
	isHelp := false
	for i, a := range args {
		// サブコマンドとしての "help"
		if i == 0 && a == "help" {
			isHelp = true
			break
		}
		// フラグとしての -h/--help
		if a == "-h" || a == "--help" {
			isHelp = true
			break
		}
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		slog.Error("app failed", "error", err)
		os.Exit(1)
	}
	if !isHelp {
		slog.Info("app terminated gracefully")
	}
}

func run(c context.Context, _ *cli.Command) error {
	var handler slog.Handler
	switch strings.ToLower(loggerFormat) {
	case "text":
		handler = slog.NewTextHandler(sharedLoggerOut, handlerOptions)
	case "json":
		handler = slog.NewJSONHandler(sharedLoggerOut, handlerOptions)
	}
	sharedLogger = slog.New(handler)

	// Profiling
	err := initProfiling()
	if err != nil {
		return fmt.Errorf("init profiling error: %+w", err)
	}

	// Tracing
	shutdown, err := initTracing()
	if err != nil {
		return fmt.Errorf("init profiling error: %+w", err)
	}
	defer shutdown()

	// OpemMetrics: Prometheus/Alloy
	err = initMetrics()
	if err != nil {
		return fmt.Errorf("init metrics error: %+w", err)
	}

	// Setup router
	router, err := setupRouter()
	if err != nil {
		return fmt.Errorf("init router error: %+w", err)
	}

	// Register APIs
	router.GET("/ping", handlePing)

	// Run server
	err = router.Run("0.0.0.0:8080")
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http server error: %+w", err)
	}

	return nil
}

// https://pkg.go.dev/github.com/grafana/pyroscope-go#Logger
// https://github.com/grafana/pyroscope-go/blob/v1.2.2/logger.go#L21
// type Logger interface {
//     Infof(_ string, _ ...interface{})
//     Debugf(_ string, _ ...interface{})
//     Errorf(_ string, _ ...interface{})
// }

type PyroscopeCustomLogger struct{}

func (p *PyroscopeCustomLogger) Infof(format string, args ...any) {

	// https://github.com/grafana/pyroscope-go/blob/v1.2.2/session.go#L80-L85
	switch {
	case format == "starting profiling session:":
		return
	case strings.HasPrefix(format, "  AppName:        "):
		format = "starting profiling session: AppName: %+v"
	case strings.HasPrefix(format, "  Tags:           "):
		format = "starting profiling session: Tags: %+v"
	case strings.HasPrefix(format, "  ProfilingTypes: "):
		format = "starting profiling session: ProfilingTypes: %+v"
	case strings.HasPrefix(format, "  DisableGCRuns:  "):
		format = "starting profiling session: DisableGCRuns: %+v"
	case strings.HasPrefix(format, "  UploadRate:     "):
		format = "starting profiling session: UploadRate: %+v"
	}

	sharedLogger.Info("pyroscope", "log", fmt.Sprintf(format, args...))
}

func (p *PyroscopeCustomLogger) Debugf(format string, args ...any) {
	sharedLogger.Debug("pyroscope", "log", fmt.Sprintf(format, args...))
}

func (p *PyroscopeCustomLogger) Errorf(format string, args ...any) {
	sharedLogger.Error("pyroscope", "log", fmt.Sprintf(format, args...))
}

// var pyroscopeLogger = pyroscope.StandardLogger
var pyroscopeLogger = &PyroscopeCustomLogger{}

func initProfiling() error {
	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	hostname, _ := os.Hostname()

	pyroscope.Start(pyroscope.Config{
		ApplicationName: appName,

		// replace this with the address of pyroscope server
		ServerAddress: pyroscopeURL,

		// you can disable logging by setting this to nil
		Logger: pyroscopeLogger,

		// you can provide static tags via a map:
		Tags: map[string]string{
			"hostname": hostname,
		},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	return nil
}

func initTracing() (func() error, error) {
	// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp#example-package

	// 内部ロガーをlog/slogに差し替え
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		sharedLogger.Error("opentelemetry export error", "error", err)
	}))

	ctx := context.Background()
	exp, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(tempoHostPort),
		otlptracehttp.WithInsecure(), // TLS なし
	)
	if err != nil {
		return nil, fmt.Errorf("otlptracehttp creation error: %+w", err)
	}
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appName),
		)),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return func() error {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("tracer_trovider shutdown error: %+w", err)
		}
		return nil
	}, nil
}

var httpRequests *prometheus.HistogramVec

func initMetrics() error {
	httpRequests = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: appName,
			Subsystem: "http_server",
			Name:      "request_duration_seconds",
			Help:      "HTTP リクエストの処理に要した時間（秒）",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
	prometheus.MustRegister(httpRequests)

	return nil
}

func setupRouter() (*gin.Engine, error) {
	// Gin
	gin.SetMode(gin.ReleaseMode)

	// https://github.com/gin-gonic/gin/blob/v1.10.0/gin.go#L224C2-L224C34
	// gin.Default()内では、engine.Use(Logger(), Recovery()) を読んでいる. gin.Logger()が先.
	// router := gin.Default()
	router := gin.New()

	// Custom logger
	// https://github.com/gin-gonic/gin/blob/v1.10.0/logger.go#L212-L281
	// https://github.com/gin-gonic/gin/blob/v1.10.0/logger.go#L196-L200
	// https://github.com/gin-gonic/gin/blob/v1.10.0/logger.go#L60
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// ↓デフォルト実装
		// https://github.com/gin-gonic/gin/blob/v1.10.0/logger.go#L141-L161

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		// log/slogで出力する
		var buf bytes.Buffer
		var handler slog.Handler
		switch strings.ToLower(loggerFormat) {
		case "text":
			handler = slog.NewTextHandler(&buf, handlerOptions)
		case "json":
			handler = slog.NewJSONHandler(&buf, handlerOptions)
		}
		logger := slog.New(handler)
		logger.Info(
			"gin access log",
			// "%v |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s"
			"access_timestamp", fmt.Sprintf("%v", param.TimeStamp.Format("2006/01/02 - 15:04:05")),
			"status_code", fmt.Sprintf("%3d", param.StatusCode),
			"latency", fmt.Sprintf("%13v", param.Latency),
			"client_ip", fmt.Sprintf("%15s", param.ClientIP),
			"method", fmt.Sprintf("%-7s", param.Method),
			"path", fmt.Sprintf("%#v", param.Path),
			"error_message", fmt.Sprintf("%s", param.ErrorMessage),
		)
		return buf.String()
	}))

	// Default recovery
	router.Use(gin.Recovery())

	// Tracing middleware
	router.Use(otelgin.Middleware(appName))

	// Prometheus middleware
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		status := fmt.Sprint(c.Writer.Status())
		httpRequests.WithLabelValues(c.FullPath(), c.Request.Method, status).Observe(duration)
	})

	// Metrics Endpoint: Prometheus/Alloy
	router.GET("/metrics/prometheus", gin.WrapH(promhttp.Handler()))

	return router, nil
}

func handlePing(c *gin.Context) {
	// Trace
	_, span := tracer.Start(c.Request.Context(), "1", oteltrace.WithAttributes(attribute.String("id", "hoge")))
	defer span.End()

	innerFunc001(c)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func innerFunc001(c *gin.Context) {
	// Trace
	_, span := tracer.Start(c.Request.Context(), "2", oteltrace.WithAttributes(attribute.String("id", "hoge")))
	defer span.End()

	innerFunc002(c)
}

func innerFunc002(c *gin.Context) {
	// Trace
	_, span := tracer.Start(c.Request.Context(), "3", oteltrace.WithAttributes(attribute.String("id", "hoge")))
	defer span.End()

	innerFunc003(c)
}

func innerFunc003(c *gin.Context) {
	// Trace
	_, span := tracer.Start(c.Request.Context(), "4", oteltrace.WithAttributes(attribute.String("id", "hoge")))
	defer span.End()
}
