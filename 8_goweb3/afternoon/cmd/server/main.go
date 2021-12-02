package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/8_goweb3/afternoon/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/8_goweb3/afternoon/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't open .env file")
	}

	router := gin.Default()
	repo := users.NewRepository()
	service := users.NewService(repo)
	hand := handler.NewUser(service)

	group := router.Group("/users")
	group.POST("", hand.Store())
	group.GET("", hand.GetAll())
	group.PUT("/:id", hand.Update())
	group.PATCH("/:id", hand.UpdateLastNameAge())
	group.DELETE("/:id", hand.Delete())

	router.Run()
}
