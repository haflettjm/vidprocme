package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func ConsoleLog(format string, args ...interface{}) {
	// Want this to be put to STDOut???
	fmt.Fprintf(gin.DefaultWriter, format, time.Now())
}
