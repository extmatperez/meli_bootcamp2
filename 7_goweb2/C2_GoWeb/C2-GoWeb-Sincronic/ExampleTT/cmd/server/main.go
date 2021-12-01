package main

import (
	users "github.com/extmatperez/meli_bootcamp2/7_goweb2/C2-GoWeb/C2-GoWeb-Sincronic/ExampleTT/internal/users"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/C2_GoWeb/C2-GoWeb-Sincronic/ExampleTT/cmd/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repo := users.NewRepository()
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	router.GET("/users/get", controller.GetAll())
	router.POST("/users/add", controller.GetAll())
	router.Run()
}
