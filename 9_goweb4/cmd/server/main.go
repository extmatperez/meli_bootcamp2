package main

import (
	"log"
	"os"

	handlers "github.com/extmatperez/meli_bootcamp2/9_goweb4/cmd/server/handler"
	middlewares "github.com/extmatperez/meli_bootcamp2/9_goweb4/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/docs"
	products "github.com/extmatperez/meli_bootcamp2/9_goweb4/internal/products"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bootcamp Go W2
// @version 1.0
// @description This is an example of API with Golang

// @host localhost:8080
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(middlewares.ValidateToken())

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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
