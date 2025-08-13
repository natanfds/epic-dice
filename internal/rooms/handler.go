package rooms

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/natanfds/epic-dice/internal/rooms/ws"
	"github.com/natanfds/epic-dice/utils"
)

type Handler struct {
	repo *RoomRepository
}

func (h *Handler) WS(c *gin.Context) {
	channelName := c.Params.ByName("room")
	if channelName == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	wsClient := ws.NewClient(conn)
	channel := ws.Hub.GetOrCreateChannel(channelName)
	channel.AddClient(wsClient)

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
		_ = json.Unmarshal(message, &msg)
		errMsg := utils.Validate.Struct(msg)

		var cmd CommandDTO
		_ = json.Unmarshal(message, &cmd)
		errCmd := utils.Validate.Struct(cmd)

		switch true {
		case errMsg == nil:
			inputProcessor = NewMessageProcessor(msg)
		case errCmd == nil:
			inputProcessor = NewCommandProcessor(cmd)
		default:
			conn.WriteMessage(websocket.TextMessage, []byte("Unprocessable input"))
			return
		}

		processResult, err := inputProcessor.Exec()

		if err != nil {
			return
		}

		conn.WriteMessage(websocket.TextMessage, processResult)

	}
}

func (h *Handler) Create(c *gin.Context) {
	var createData CreateRoomDTO
	bodyData := c.ShouldBind(&createData)
	validationErr := utils.Validate.Struct(createData)
	if bodyData != nil || validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.repo.Create(RoomModel{Name: createData.Name}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room created successfully"})
}

func NewRoomHandler(repo *RoomRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}
