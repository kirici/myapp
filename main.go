package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func handleHello(c *gin.Context) {
	// reqsByCode.WithLabelValues(strconv.Itoa(http.StatusOK)).Inc()
	c.JSON(200, "Hello world!")
}

func handleMath(c *gin.Context) {
	res := rand.Float64() * rand.Float64()
	if res > 0.8 {
		c.JSON(501, gin.H{"values over 0.8 not yet supported": res})
		return
	}
	c.JSON(200, res)
}

func handleWork(c *gin.Context) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
	c.JSON(200, "Done!")
}

func main() {
	router := gin.Default()
	prometheus.MustRegister(reqsByCode, errsByCode, reqDuration) // Alternatively move to init()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// Attach to all routes except prometheus itself
	router.Use(prometheusMiddleware())
	router.GET("/", handleHello)
	router.GET("/math", handleMath)
	router.GET("/work", handleWork)

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
