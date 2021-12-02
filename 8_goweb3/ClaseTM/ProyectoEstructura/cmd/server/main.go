package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/8_goweb3/ClaseTM/ProyectoEstructura/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/8_goweb3/ClaseTM/ProyectoEstructura/internal/producto"
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
		groupProducts.PUT("/:id", controller.Update())
		groupProducts.DELETE("/:id", controller.Detele())
		groupProducts.PATCH("/:id", controller.UpdateNameAndPrice())

	}

	router.Run()
}
