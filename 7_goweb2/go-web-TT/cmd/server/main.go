package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/7_goweb2/go-web-TT/cmd/server/handler"
	product "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/7_goweb2/go-web-TT/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := product.NewRepository()
	service := product.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/product/get", controller.GetAll())
	router.POST("/product/add", controller.Store())

	router.Run()
}
