package view

import (
	"dlp-ui/sidecar/ytdlp"
	"dlp-ui/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Download(router *gin.Engine) {
	var downloads sync.Map

	router.GET("/api/download", func(context *gin.Context) {
		// middlewares
		logger := context.MustGet("logger").(*logrus.Entry)

		// get all downloads
		downloads := utils.Stom(&downloads)
		logger.Debugf("got %d downloads", len(downloads))

		// success
		context.JSON(http.StatusOK, downloads)
	})

	router.POST("/api/download", func(context *gin.Context) {
		// middlewares
		logger := context.MustGet("logger").(*logrus.Entry)
		// data
		var data struct {
			URL    string `json:"url"`
			Format string `json:"format"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			// failure
			logger.Warningf("failed to get data: %v", err)
			return
		}

		// check if the url already exists and has already stopped downloading files
		value, ok := downloads.Load(data.URL)
		if ok {
			download, ok := value.(ytdlp.Download)
			if !ok {
				// failure
				logger.Error("failed to check if the url has already stopped downloading files")
				context.Status(http.StatusInternalServerError)
				return
			}

			// check if the url has already stopped downloading files
			if !download.Progress.Done {
				// failure
				logger.Warning("cannot download files: the url has not stopped downloading files")
				context.Status(http.StatusConflict)
				return
			}
		}

		// initalize a download
		downloads.Store(data.URL, ytdlp.Download{
			Title: "downloading",
			Progress: ytdlp.Progress{
				Current: -1,
				Total:   -1,
				Speed:   -1,
				Done:    false,
			},
			Errors: []string{},
		})

		// create a new downloader
		downloader, err := ytdlp.Downloader(data.URL, data.Format, &downloads)
		if err != nil {
			// failure
			logger.Errorf("failed to create a downloader: %v", err)
			context.String(http.StatusInternalServerError, "%v", err.Error())
			return
		}

		// start the downloader
		go downloader(logger)
		logger.Debug("started a downloader")

		// success
		context.Status(http.StatusOK)
	})
}
