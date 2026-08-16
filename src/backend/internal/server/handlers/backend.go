package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
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
