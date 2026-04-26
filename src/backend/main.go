package main

import (
	"dlp-ui/cmd"
	"dlp-ui/internal/web"
	"dlp-ui/internal/web/views"
	"dlp-ui/pkg/log"
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed embed/*
var embedFS embed.FS

func main() {
	level := cmd.GetLogLevel()
	logger, err := log.New("dlp-ui", level)
	if err != nil {
		panic(err)
	}

	webUI, err := fs.Sub(embedFS, "embed/webui")
	if err != nil {
		logger.Fatalf("failed to load webui: %v", err)
	}

	cmd.PrintBanner()

	mode := cmd.GetMode()
	gin.SetMode(mode)

	router := web.New(logger)

	if mode != gin.ReleaseMode {
		router.Use(cors.Default())
	}

	views.RouteParse(router)
	views.RouteDownload(router)
	router.StaticFS("/webui", http.FS(webUI))

	router.Run("localhost:5000")
}
