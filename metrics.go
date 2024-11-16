package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	reqsByCode = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of successful GET requests by HTTP code.",
		},
		[]string{"code"},
	)
	errsByCode = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of returned error responses (400 or higher), by code.",
		},
		[]string{"code"},
	)
	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response latency for handler in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
)

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// Calculate request t and observe in the histogram
		t := time.Since(start).Seconds()
		reqDuration.WithLabelValues(c.FullPath()).Observe(t)

		// Record traffic and error metrics
		status := c.Writer.Status()
		reqsByCode.WithLabelValues(strconv.Itoa(status)).Inc()
		if status >= 400 {
			errsByCode.WithLabelValues(strconv.Itoa(status)).Inc()
		}
	}
}
