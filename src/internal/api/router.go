package router

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"vidprocme/internal/config"
	"vidprocme/internal/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.GET("/greet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Success %v-enviornment is running", cfg.EnvType),
			"time":    time.Now().Format(time.RFC3339),
		})
		utils.ConsoleLog("Hello From the Console! %s\n", time.Now()) // this should work??
	})

	return router
}

func RunRouter(cfg *config.Config, router *gin.Engine) {
	err := router.Run(":" + strconv.Itoa(cfg.Port))
	if err != nil {
		utils.ConsoleLog("Error running router: %v\n", err)
	}
}

func StartServer(cfg *config.Config) {
	router := InitRouter(cfg)
	RunRouter(cfg, router)
}

func StopServer() {
	// Implement server stop logic here
}

func ShutdownServer() {
	// Implement server shutdown logic here
}

func RestartServer(cfg *config.Config) {
	StopServer()
	StartServer(cfg)
}
