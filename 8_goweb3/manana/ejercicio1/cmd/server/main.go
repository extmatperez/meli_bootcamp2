package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/8_goweb3/tarde/ejercicio1/cmd/server/handler"

	users "github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/8_goweb3/manana/ejercicio1/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := users.NewRepository()
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("/usuarios/:id", controller.Update())
	router.PATCH("/usuarios/:id", controller.UpdateNombre())
	router.DELETE("/usuarios/:id", controller.Delete())

	router.Run()
}
