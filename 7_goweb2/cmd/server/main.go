package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/castiglione_adrian/7_goweb2/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/castiglione_adrian/7_goweb2/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repository := producto.NewRepository()
	service := producto.NewService(repository)
	controller := handler.NewProducto(service)

	router.GET("/productos", controller.GetAll())
	router.GET("/productos/:id", controller.Get())
	router.POST("/productos", controller.AddOne())

	router.Run()
}
