package views

import (
	"dlp-ui/internal/ytdlp"
	"dlp-ui/pkg/utils"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var downloads = []ytdlp.Download{}

func RouteDownload(router *gin.Engine) {
	router.GET("/api/download", func(context *gin.Context) {
		context.JSON(http.StatusOK, downloads)
	})

	router.POST("/api/download", func(context *gin.Context) {
		logger := context.MustGet("logger").(*logrus.Entry)

		browser, err := utils.ParseBrowser(context.Request.UserAgent())
		if err != nil {
			logger.Errorf("failed to parse user agent: %v", err)
			context.String(http.StatusInternalServerError, "%v", err)
			return
		}

		var data struct {
			URL    string `json:"url"`
			Format string `json:"format"`
		}
		err = context.BindJSON(&data)
		if err != nil {
			logger.Warningf("failed to get data: %v", err)
			return
		}

		exists := slices.ContainsFunc(downloads, func(download ytdlp.Download) bool {
			return data.URL == download.URL
		})
		if exists {
			index := slices.IndexFunc(downloads, func(download ytdlp.Download) bool {
				return data.URL == download.URL
			})
			download := downloads[index]
			if download.Task.Done {
				downloads = slices.DeleteFunc(downloads, func(download ytdlp.Download) bool {
					return data.URL == download.URL
				})
			} else {
				logger.Warning("cannot create a new downloader: repeated requests for downloading")
				context.Status(http.StatusConflict)
				return
			}
		}

		downloads = append(downloads, ytdlp.Download{URL: data.URL})

		downloader, err := ytdlp.NewDownloader(browser, data.URL, data.Format, downloads)
		if err != nil {
			logger.Errorf("failed to create a new downloader: %v", err.Error())
			context.String(http.StatusInternalServerError, "%v", err)
			return
		}

		go downloader(logger)

		context.Status(http.StatusOK)
	})

	router.DELETE("/api/download", func(context *gin.Context) {
		logger := context.MustGet("logger").(*logrus.Entry)

		var data struct {
			URL string `json:"url"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			logger.Warningf("failed to get data: %v", err)
			return
		}

		downloads = slices.DeleteFunc(downloads, func(download ytdlp.Download) bool {
			return data.URL == download.URL
		})

		context.Status(http.StatusOK)
	})
}
