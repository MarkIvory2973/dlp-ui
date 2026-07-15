package handlers

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFrontend(router *gin.Engine, embedFS embed.FS) error {
	frontendFS, err := fs.Sub(embedFS, "embed/frontend")
	if err != nil {
		return err
	}

	router.StaticFS("/frontend", http.FS(frontendFS))

	return nil
}
