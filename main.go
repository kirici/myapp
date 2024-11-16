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
