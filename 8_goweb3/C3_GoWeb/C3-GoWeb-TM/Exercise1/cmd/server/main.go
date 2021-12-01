package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-TM/Exercise1/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-TM/Exercise1/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repo := users.NewRepository()
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	router.GET("/users/get", controller.GetAll())
	router.POST("/users/add", controller.Store())
	router.PUT("/users/:id", controller.Update())
	router.Run()
}
