package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	// OpenMetrics
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

var tracer = otel.Tracer("gin-server")

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Tracing
	// https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp#example-package
	ctx := context.Background()
	exp, err := otlptracehttp.New(
		ctx,
		// otlptracehttp.WithEndpoint("tempo:4318"), // ポート 4318 を指定
		otlptracehttp.WithEndpoint("host.docker.internal:4318"), // ポート 4318 を指定
		otlptracehttp.WithInsecure(),                            // TLS なし
	)
	if err != nil {
		panic(err)
	}
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-gin-service"),
		)),
	)
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	// OpemMetrics: Prometheus/Alloy
	httpRequests := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "myapp",
			Subsystem: "http_server",
			Name:      "request_duration_seconds",
			Help:      "HTTP リクエストの処理に要した時間（秒）",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
	prometheus.MustRegister(httpRequests)

	// Gin
	router := gin.Default()

	// Tracing middleware
	router.Use(otelgin.Middleware("my-gin-service"))

	// Prometheus middleware
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		status := fmt.Sprint(c.Writer.Status())
		httpRequests.WithLabelValues(c.FullPath(), c.Request.Method, status).
			Observe(duration)
	})

	// Metrics Endpoint: Prometheus/Alloy
	router.GET("/metrics/prometheus", gin.WrapH(promhttp.Handler()))

	// API
	router.GET("/ping", func(c *gin.Context) {
		// Trace
		_, span := tracer.Start(c.Request.Context(), "1", oteltrace.WithAttributes(attribute.String("id", "z")))
		defer span.End()

		innerFunc(c)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Run server
	router.Run("0.0.0.0:8080")
}

func innerFunc(c *gin.Context) {
	// Trace
	_, span := tracer.Start(c.Request.Context(), "2", oteltrace.WithAttributes(attribute.String("id", "z")))
	defer span.End()
}
