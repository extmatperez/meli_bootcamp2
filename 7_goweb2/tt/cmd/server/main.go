package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/7_goweb2/tt/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/7_goweb2/tt/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("getAll", controller.GetAll())
	router.POST("store", controller.Store())

	router.Run()
}
