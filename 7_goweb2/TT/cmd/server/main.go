package main

import (
	handlers "github.com/extmatperez/meli_bootcamp2/tree/ziliotto_matias/7_goweb2/TT/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/tree/ziliotto_matias/7_goweb2/TT/internal/products"
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
		products.GET("/:id", productsController.FindById())
		products.POST("/", productsController.Store())
	}

	router.Run()
}
