package handlers

import (
	"dlp-ui/internal/ytdlp"
	"net/http"

	"github.com/gin-gonic/gin"
)

var downloads = ytdlp.Downloads{}

func HandleDownloads(router *gin.Engine) {
	router.GET("/api/download", func(context *gin.Context) {
		context.JSON(http.StatusOK, downloads)
	})

	router.POST("/api/download", func(context *gin.Context) {
		var data struct {
			URL    string `json:"url"`
			Format string `json:"format"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			return
		}

		if downloads.Contains(data.URL) {
			index := downloads.Index(data.URL)
			if !downloads[index].Job.Done {
				context.Status(http.StatusConflict)
				return
			}

			downloads.Delete(data.URL)
		}

		downloads.Append(data.URL)

		downloader, err := ytdlp.NewDownloader(data.URL, data.Format)
		if err != nil {
			context.String(http.StatusInternalServerError, "%v", err)
			return
		}

		go downloader(downloads)

		context.Status(http.StatusOK)
	})

	router.DELETE("/api/download", func(context *gin.Context) {
		var data struct {
			URL string `json:"url"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			return
		}

		if !downloads.Contains(data.URL) {
			context.Status(http.StatusNotFound)
			return
		}

		index := downloads.Index(data.URL)
		if !downloads[index].Job.Done {
			context.Status(http.StatusLocked)
			return
		}

		downloads.Delete(data.URL)

		context.Status(http.StatusOK)
	})
}
