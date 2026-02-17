package view

import "github.com/gin-gonic/gin"

func UI(router *gin.Engine) {
	router.Static("/ui", "ui")
}
