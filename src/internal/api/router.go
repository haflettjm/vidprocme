package router

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"vidprocme/internal/config"
	"vidprocme/internal/queue"
	"vidprocme/internal/scheduler"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(cfg *config.Config, logger *zap.Logger, sched *scheduler.Scheduler, queue *queue.Queue) *gin.Engine {
	router := gin.Default()
	router.GET("/greet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Success %v-enviornment is running", cfg.EnvType),
			"time":    time.Now().Format(time.RFC3339),
		})
		logger.Info("Hello From the Console!", zap.String("time", time.Now().Format(time.RFC3339)))
	})

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"time":   time.Now().Format(time.RFC3339),
		})
		logger.Info("Health check successful", zap.String("time", time.Now().Format(time.RFC3339)))
	})

	return router
}

func RunRouter(cfg *config.Config, router *gin.Engine, logger *zap.Logger, sched *scheduler.Scheduler, queue *queue.Queue) {
	err := router.Run(":" + strconv.Itoa(cfg.Port))
	if err != nil {
		logger.Error("Error running router", zap.Error(err))
	}
}

func StartServer(cfg *config.Config, logger *zap.Logger, sched *scheduler.Scheduler, queue *queue.Queue) {
	router := InitRouter(cfg, logger, sched, queue)
	RunRouter(cfg, router, logger, sched, queue)
}

func StopServer(logger *zap.Logger) {
	// Implement server stop logic here
	logger.Info("Server stopped")
	os.Exit(0)
}

func ShutdownServer() {
	// Implement server shutdown logic here
}

func RestartServer(cfg *config.Config, logger *zap.Logger, sched *scheduler.Scheduler, queue *queue.Queue) {
	StopServer(logger)
	StartServer(cfg, logger, sched, queue)
}
