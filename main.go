package main

import (
	"github.com/gin-gonic/gin"

	"github.com/natanfds/epic-dice/internal/ping"
	"github.com/natanfds/epic-dice/internal/rooms"
)

func main() {
	router := gin.Default()

	router.GET("/ping", ping.Handler)

	router.GET("/room/*room", rooms.Handler)

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
