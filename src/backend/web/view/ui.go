package view

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed ui/*
var ui embed.FS

func UI(router *gin.Engine) {
	router.StaticFS("/ui", http.FS(ui))
}
