package main

import (
	"dlp-ui/cmd"
	"dlp-ui/internal/server"
	"dlp-ui/internal/server/handlers"
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
)

//go:embed embed/*
var embedFS embed.FS

func StartServer() {
	var mode string
	if cmd.GetDebug() {
		mode = gin.DebugMode
	} else {
		mode = gin.ReleaseMode
	}

	router := server.New(mode)

	handlers.HandleParse(router)
	handlers.HandleDownload(router)
	handlers.HandleBackend(router)
	handlers.HandleFrontend(router, embedFS)

	router.Run("localhost:5000")
}

func StartURL() {
	err := browser.OpenURL("http://localhost:5000/frontend")
	if err != nil {
		panic(err)
	}
}

func main() {
	go StartServer()
	go StartURL()
	select {}
}
