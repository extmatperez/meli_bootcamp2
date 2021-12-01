package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := producto.NewRepository()
	service := producto.NewService(repo)
	controller := handler.NewProducto(service)

	router.GET("/productos", controller.GetAll())
	router.POST("/addProductos", controller.Store())
	router.PUT("/modify/:id", controller.Modify())
	router.PATCH("/modifyNaPr/:id", controller.ModifyNamePrice())
	router.DELETE("/delete/:id", controller.Delete())

	router.Run()
}
