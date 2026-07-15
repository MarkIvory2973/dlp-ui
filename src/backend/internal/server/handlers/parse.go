package handlers

import (
	"dlp-ui/internal/ytdlp"
	"net/http"

	"github.com/gin-gonic/gin"
)

var parseds = ytdlp.Parseds{}

func HandleParse(router *gin.Engine) {
	router.GET("/api/parse", func(context *gin.Context) {
		context.JSON(http.StatusOK, parseds)
	})

	router.POST("/api/parse", func(context *gin.Context) {
		var data struct {
			URL string `json:"url"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			return
		}

		if parseds.Contains(data.URL) {
			context.Status(http.StatusConflict)
			return
		}

		parseds.Append(data.URL)

		parser, err := ytdlp.NewParser(data.URL)
		if err != nil {
			context.String(http.StatusInternalServerError, "%v", err)
			return
		}

		go parser(parseds)

		context.Status(http.StatusOK)
	})

	router.DELETE("/api/parse", func(context *gin.Context) {
		var data struct {
			URL string `json:"url"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			return
		}

		if !parseds.Contains(data.URL) {
			context.Status(http.StatusNotFound)
			return
		}

		index := parseds.Index(data.URL)
		if !parseds[index].Job.Done {
			context.Status(http.StatusLocked)
			return
		}

		parseds.Delete(data.URL)

		context.Status(http.StatusOK)
	})
}
