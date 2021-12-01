package main

import (
	"github.com/extmatperez/meli_bootcamp2/pecora_estefania/8_goweb3/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/pecora_estefania/8_goweb3/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {

	/////////////////////// start server /////////////////////////////
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	repo.ReadData()

	products := router.Group("/products")

	products.GET("/", controller.GetAll())
	//products.GET("/:id", getProductbyID)
	// products.GET("/products/filter/select", getbyFilter)

	products.POST("/addproduct", controller.AddProduct())

	router.Run()

}
