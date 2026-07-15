package main

import (
	"dlp-ui/internal/server"
	"dlp-ui/internal/server/handlers"
	"embed"

	"github.com/pkg/browser"
)

//go:embed embed/*
var embedFS embed.FS

func StartServer() {
	router := server.New()

	handlers.HandleParse(router)
	handlers.HandleDownloads(router)
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
