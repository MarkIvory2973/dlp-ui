package view

import (
	"dlp-ui/sidecar/ytdlp"
	"dlp-ui/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Parse(router *gin.Engine) {
	var parseds sync.Map

	router.GET("/api/parse", func(context *gin.Context) {
		// middlewares
		logger := context.MustGet("logger").(*logrus.Entry)

		// get all parseds
		parseds := utils.Stom(&parseds)
		logger.Debugf("got %d parseds", len(parseds))

		// success
		context.JSON(http.StatusOK, parseds)
	})

	router.POST("/api/parse", func(context *gin.Context) {
		// middlewares
		logger := context.MustGet("logger").(*logrus.Entry)
		// data
		var data struct {
			URL string `json:"url"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			// failure
			logger.Warningf("failed to get data: %v", err)
			return
		}

		// check if the url already exists
		_, ok := parseds.Load(data.URL)
		if ok {
			// failure
			logger.Warningf("cannot create a new parser: the url already exists")
			context.Status(http.StatusConflict)
			return
		}

		// initalize the parsed
		parseds.Store(data.URL, ytdlp.Parsed{
			Entries: []map[string]any{
				{
					"title":   "parsing",
					"formats": []string{},
				},
			},
			Errors: []string{},
		})

		// create a new parser
		parser, err := ytdlp.Parser(data.URL, &parseds)
		if err != nil {
			// failure
			logger.Errorf("failed to create a new parser: %v", err)
			context.String(http.StatusInternalServerError, "%v", err.Error())
			return
		}

		// start the parser
		go parser(logger)
		logger.Debugf("started a parser")

		// success
		context.Status(http.StatusOK)
	})

	router.DELETE("/api/parse", func(context *gin.Context) {
		// middlewares
		logger := context.MustGet("logger").(*logrus.Entry)
		// data
		var data struct {
			URL string `json:"url"`
		}
		err := context.BindJSON(&data)
		if err != nil {
			// failure
			logger.Warningf("failed to get data: %v", err)
			return
		}

		// delete a parsed
		parseds.Delete(data.URL)
		logger.Debug("deleted 1 parsed")

		// success
		context.Status(http.StatusOK)
	})
}
