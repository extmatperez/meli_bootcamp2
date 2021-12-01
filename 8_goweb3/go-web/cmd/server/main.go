package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web/cmd/server/handler"
	product "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := product.NewRepository()
	service := product.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/product/get", controller.GetAll())
	router.POST("/product/add", controller.Store())
	router.PUT("/product/put/:id", controller.Update())
	router.PATCH("/product/patch/:id", controller.UpdateNombre())
	router.PATCH("/product/patchprecio/:id", controller.UpdatePrecio())
	router.DELETE("/product/delete/:id", controller.Delete())
	router.Run()
}
