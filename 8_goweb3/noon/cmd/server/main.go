package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/8_goweb3/noon/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/8_goweb3/noon/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
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
