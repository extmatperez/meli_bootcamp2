package main

import (
	"log"

	handlers "github.com/extmatperez/meli_bootcamp2/9_goweb4/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/9_goweb4/internal/products"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	dbStore := store.New(store.FileType, "../../internal/products/products.json")
	productsRepository := products.NewRepository(dbStore)
	productsService := products.NewService(productsRepository)
	productsController := handlers.NewProduct(productsService)

	products := router.Group("/products")
	{
		products.GET("/", productsController.GetAll())
		products.GET("/:id", productsController.FindById())
		products.POST("/", productsController.Store())
		products.PUT("/:id", productsController.Update())
		products.DELETE("/:id", productsController.Delete())
		products.PATCH("/:id", productsController.UpdateNameAndPrice())
	}

	router.Run()
}
