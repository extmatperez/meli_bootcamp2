package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/internal/products"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el ambiente")
	}

	router := gin.Default()

	db := store.New("file", os.Getenv("FILEPATH"))
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	controller := handler.NewProduct(service)

	productsRoute := router.Group("products")
	productsRoute.GET("", controller.GetAll())
	productsRoute.GET("/filter", controller.Filter())
	productsRoute.GET("/:id", controller.FindById())
	productsRoute.POST("", controller.Store())
	productsRoute.PUT("/:id", controller.Update())
	productsRoute.DELETE("/:id", controller.Delete())
	productsRoute.PATCH("/:id", controller.UpdateNameAndPrice())

	router.Run()
}
