package main

import (
	handlers "github.com/extmatperez/meli_bootcamp2/tree/ziliotto_matias/8_goweb3/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/tree/ziliotto_matias/8_goweb3/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	productsRepository := products.NewRepository()
	productsService := products.NewService(productsRepository)
	productsController := handlers.NewProduct(productsService)

	products := router.Group("/products")
	{
		products.GET("/", productsController.GetAll())
		products.GET("/load", productsController.LoadProducts())
		products.GET("/:id", productsController.FindById())
		products.POST("/", productsController.Store())
		products.PUT("/:id", productsController.Update())
		products.DELETE("/:id", productsController.Delete())
	}

	router.Run()
}
