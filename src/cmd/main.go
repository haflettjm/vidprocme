package main

import (
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
	cfg.Load()

}
