package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/natanfds/epic-dice/internal/database"
	"github.com/natanfds/epic-dice/internal/ping"
	"github.com/natanfds/epic-dice/internal/rooms"

	_ "github.com/natanfds/epic-dice/docs"
)

// @title Epic Dice API
// @version 1.0
// @description API para jogar TTRPG
// @host localhost:8080
func startAPI(port string) error {
	router := gin.Default()

	router.Use(
		gin.Logger(),
		gin.Recovery(), // Transforma os panics em 500
	)
	router.GET("/ping", ping.Handler)

	{
		db, err := database.CreateSQLDB(
			rooms.RoomModel{},
		)
		if err != nil {
			return err
		}

		v1 := router.Group("/v1")

		roomHandler := rooms.NewRoomHandler(rooms.NewRoomRepository(db))
		v1.GET("/room/*room", roomHandler.WS)
		v1.POST("/room", roomHandler.Create)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":" + port)
	return err
}

func main() {
	if err := startAPI("8080"); err != nil {
		panic(err)
	}
}
