package main

import (
	"fmt"
	"time"

	"vidprocme/internal/config"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Ref *gin.Engine
}

type Handler struct{}

func ConsoleLog(format string, args ...interface{}) {
	// Want this to be put to STDOut???
	fmt.Fprintf(gin.DefaultWriter, format, time.Now())
}

func main() {
	cfg := &config.Config{}
	cfg.Load()
	router := gin.Default()
	router.GET("/greet", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
			"time":    time.Now().Format(time.RFC3339),
		})
		ConsoleLog("Hello From the Console! %s\n", time.Now()) // this should work??
	})

	router.Run()

}
