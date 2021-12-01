package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-Sincronic/ExampleTM/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-Sincronic/ExampleTM/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repo := users.NewRepository()
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	router.GET("/users/get", controller.GetAll())
	router.POST("/users/add", controller.Store())
	router.Run()
}
