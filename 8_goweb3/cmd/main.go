package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {

	/////////////////////// start server /////////////////////////////
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	productsroute := router.Group("/products")

	productsroute.GET("/", controller.GetAll())
	//products.GET("/:id", getProductbyID)
	// products.GET("/products/filter/select", getbyFilter)

	productsroute.POST("/addproduct", controller.AddProduct())
	productsroute.PUT("/updateproduct/:id", controller.UpdateProduct())

	router.Run()

}
