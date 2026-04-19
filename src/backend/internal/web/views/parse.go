package views

import (
	"dlp-ui/internal/ytdlp"
	"dlp-ui/pkg/utils"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var parseds = []ytdlp.Parsed{}

func RouteParse(router *gin.Engine) {
	router.GET("/api/parse", func(context *gin.Context) {
		context.JSON(http.StatusOK, parseds)
	})

	router.POST("/api/parse", func(context *gin.Context) {
		logger := context.MustGet("logger").(*logrus.Entry)

		browser, err := utils.ParseBrowser(context.Request.UserAgent())
		if err != nil {
			logger.Errorf("failed to parse user agent: %v", err)
			context.String(http.StatusInternalServerError, "%v", err)
			return
		}

		var data struct {
			URL string `json:"url"`
		}
		err = context.BindJSON(&data)
		if err != nil {
			logger.Warningf("failed to get data: %v", err)
			return
		}

		exists := slices.ContainsFunc(parseds, func(parsed ytdlp.Parsed) bool {
			return data.URL == parsed.URL
		})
		if exists {
			logger.Warning("failed to create a new parser: repeated requests for parsing")
			context.Status(http.StatusConflict)
			return
		}

		parseds = append(parseds, ytdlp.Parsed{URL: data.URL})

		parser, err := ytdlp.NewParser(browser, data.URL, parseds)
		if err != nil {
			logger.Errorf("failed to create a new parser: %v", err)
			context.String(http.StatusInternalServerError, "%v", err)
			return
		}

		go parser(logger)

		context.Status(http.StatusOK)
	})

	router.DELETE("/api/parse", func(context *gin.Context) {
		logger := context.MustGet("logger").(*logrus.Entry)

		var data struct {
			URL string `json:"url"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			logger.Warningf("failed to get data: %v", err)
			return
		}

		parseds = slices.DeleteFunc(parseds, func(parsed ytdlp.Parsed) bool {
			return data.URL == parsed.URL
		})

		context.Status(http.StatusOK)
	})
}
