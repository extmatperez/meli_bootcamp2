package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/7_goweb2/tarde/ejercicio1/cmd/server/handler"

	users "github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/7_goweb2/tarde/ejercicio1/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := users.NewRepository()
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	router.GET("/product/get", controller.GetAll())
	router.POST("/product/add", controller.Store())

	router.Run()
}
