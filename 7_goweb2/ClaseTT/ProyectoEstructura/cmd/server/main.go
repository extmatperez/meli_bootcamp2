package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/7_goweb2/ClaseTT/ProyectoEstructura/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/7_goweb2/ClaseTT/ProyectoEstructura/internal/producto"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := producto.NewRepository()
	service := producto.NewService(repo)
	controller := handler.NewProductoController(service)

	groupProducts := router.Group("api/productos")
	{
		groupProducts.GET("/", controller.GetAll())
		groupProducts.POST("/", controller.Store())
	}

	router.Run()
}
