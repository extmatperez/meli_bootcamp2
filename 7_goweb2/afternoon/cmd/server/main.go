package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/7_goweb2/afternoon/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/7_goweb2/afternoon/internal/users"
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

	router.Run()
}
