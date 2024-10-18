package main

import (
	"log"
	"math/rand"
	"os"
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

func handleHello(c *gin.Context) {
	// reqsByCode.WithLabelValues(strconv.Itoa(http.StatusOK)).Inc()
	reqsByCode.WithLabelValues("200").Inc()
	c.JSON(200, "Hello world!")
}

func handleMath(c *gin.Context) {
	reqsByCode.WithLabelValues("200").Inc()
	// c.JSON(200, gin.H{"result": rand.Float64() * rand.Float64()})
	c.JSON(200, rand.Float64()*rand.Float64())
}

func handleWork(c *gin.Context) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	reqsByCode.WithLabelValues("200").Inc()
	c.JSON(200, "Done!")
}

func main() {
	prometheus.MustRegister(reqsByCode) // Alternatively move to init()
	router := gin.Default()
	router.GET("/", handleHello)
	router.GET("/math", handleMath)
	router.GET("/work", handleWork)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
