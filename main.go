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

var reqsByCode = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of GET requests by HTTP code.",
	},
	[]string{"code"},
)

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := c.Writer.Status()
		reqsByCode.WithLabelValues(strconv.Itoa(status)).Inc()
	}
}

func handleHello(c *gin.Context) {
	// reqsByCode.WithLabelValues(strconv.Itoa(http.StatusOK)).Inc()
	c.JSON(200, "Hello world!")
}

func handleMath(c *gin.Context) {
	// c.JSON(200, gin.H{"result": rand.Float64() * rand.Float64()})
	c.JSON(200, rand.Float64()*rand.Float64())
}

func handleWork(c *gin.Context) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	c.JSON(200, "Done!")
}

func main() {
	router := gin.Default()
	prometheus.MustRegister(reqsByCode) // Alternatively move to init()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// Attach to all routes except prometheus itself
	router.Use(prometheusMiddleware())
	router.GET("/", handleHello)
	router.GET("/math", handleMath)
	router.GET("/work", handleWork)

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
