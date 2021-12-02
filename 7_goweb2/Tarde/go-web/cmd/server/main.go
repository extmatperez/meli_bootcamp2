package main

import (
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repository := products.NewRepository()
	service := products.NewService(repository)
	controller := handler.NewProduct(service)

	productsRoute := router.Group("products")
	productsRoute.GET("", controller.GetAll())
	//productsRoute.GET("/filter", Filter)
	productsRoute.GET("/:id", controller.FindById())
	productsRoute.POST("", controller.Store())
	productsRoute.PUT("/:id", controller.Update())
	productsRoute.DELETE("/:id", controller.Delete())
	productsRoute.PATCH("/:id", controller.UpdateNameAndPrice())

	router.Run()
}
