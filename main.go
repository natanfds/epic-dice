package main

import (
	"github.com/gin-gonic/gin"

	"github.com/natanfds/epic-dice/internal/ping"
)

func main() {
	router := gin.Default()

	router.GET("/ping", ping.Handler)

	router.Run()
}
