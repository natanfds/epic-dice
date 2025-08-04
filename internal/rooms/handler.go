package rooms

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	defer conn.Close()

	for {
		time.Sleep(time.Second)
		var inputProcessor InputProcessor
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Error reading message"))
			return
		}

		var msg MessageDTO
		errMsg := json.Unmarshal(message, &msg)

		var cmd CommandDTO
		errCmd := json.Unmarshal(message, &cmd)

		switch true {
		case errMsg == nil:
			inputProcessor = NewMessageProcessor(msg)
		case errCmd == nil:
			inputProcessor = NewCommandProcessor(cmd)
		default:
			return
		}

		processResult, err := inputProcessor.Exec()

		if err != nil {
			return
		}

		conn.WriteMessage(websocket.TextMessage, processResult)

	}
}
