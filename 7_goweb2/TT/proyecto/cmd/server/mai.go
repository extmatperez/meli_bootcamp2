package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/arguello_nico/7_goweb2/TT/proyecto/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/arguello_nico/7_goweb2/TT/proyecto/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := producto.NewRepository()
	service := producto.NewService(repo)
	controller := handler.NewProducto(service)

	router.GET("/products", controller.GetAll())
	router.POST("/addProducts", controller.Store())

	router.Run()
}
