package handlers

import (
	"dlp-ui/cmd"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		if cmd.GetDebug() {
			return true
		}

		content := request.Header.Get("Origin")
		origin, err := url.Parse(content)
		if err != nil {
			return false
		}

		if request.Host != origin.Host {
			return false
		}

		return true
	},
}

var parseSignal = make(chan string)
var downloadSignal = make(chan string)

func HandleBackend(router *gin.Engine) {
	router.GET("/backend/parse", func(context *gin.Context) {
		connection, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			return
		}
		defer connection.Close()

		for signal := range parseSignal {
			connection.WriteMessage(websocket.TextMessage, []byte(signal))
		}
	})

	router.GET("/backend/download", func(context *gin.Context) {
		connection, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			return
		}
		defer connection.Close()

		for signal := range downloadSignal {
			connection.WriteMessage(websocket.TextMessage, []byte(signal))
		}
	})
}
