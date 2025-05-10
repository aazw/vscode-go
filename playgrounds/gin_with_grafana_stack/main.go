package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
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

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		status := fmt.Sprint(c.Writer.Status())
		httpRequests.WithLabelValues(c.FullPath(), c.Request.Method, status).
			Observe(duration)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/metrics/prometheus", gin.WrapH(promhttp.Handler()))

	router.Run("0.0.0.0:8080")
}
