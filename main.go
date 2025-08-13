package main

import (
	"github.com/gin-gonic/gin"

	"github.com/natanfds/epic-dice/internal/ping"
	"github.com/natanfds/epic-dice/internal/rooms"
)

func main() {
	router := gin.Default()

	router.Use(
		gin.Logger(),
		gin.Recovery(), // Transforma os panics em 500
	)
	router.GET("/ping", ping.Handler)

	{
		v1 := router.Group("/v1")
		v1.GET("/room/*room", rooms.HandlerWS)
		//v1.POST("/room", rooms.HandlerCreate)
	}

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
