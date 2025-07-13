package main

import (
	"fmt"
	"os"
	router "vidprocme/internal/api"
	"vidprocme/internal/config"
	"vidprocme/internal/queue"
	"vidprocme/internal/scheduler"

	"go.uber.org/zap"
)

func main() {
	// Initialize the configuration
	cfg := &config.Config{}
	if err := cfg.Load(); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	logger := zap.NewExample()
	q := queue.New(logger, cfg)
	sched := scheduler.New(logger, q, cfg)

	fmt.Println("Server started")
	router.StartServer(cfg, logger, sched, q)

}
