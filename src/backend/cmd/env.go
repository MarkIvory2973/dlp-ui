package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GetMode() string {
	mode := os.Getenv("MODE")

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
