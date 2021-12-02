package main

import (
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/project/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/project/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	// router.PUT("/products/:id", controller.Update())
	// router.PATCH("/products/:id", controller.UpdateProd())
	// router.DELETE("/products/:id", controller.Delete())
	router.Run()

}
