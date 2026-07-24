package cmd

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMode() string {
	mode := os.Getenv("MODE")
	mode = strings.ToLower(mode)

	switch mode {
	case "debug":
		return gin.DebugMode
	case "release":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.ReleaseMode
	}
}
