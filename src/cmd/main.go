package main

import (
	"fmt"
	"os"
	router "vidprocme/internal/api"
	"vidprocme/internal/config"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Ref *gin.Engine
}

type Handler struct{}

func main() {
	// Initialize the configuration
	cfg := &config.Config{}
	if err := cfg.Load(); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}
	fmt.Println("Server started")
	router.StartServer(cfg)
}
