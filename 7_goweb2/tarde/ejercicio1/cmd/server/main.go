package main

import "github.com/gin-gonic/gin"
			"github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/7_goweb2/tarde/ejercicio1/cmd/server/handler"
			user "github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/7_goweb2/tarde/ejercicio1/internal/users"

func main() {
	router := gin.Default()

	repo := user.NewRepository()
	service := user.NewService(repo)
	controller := user.NewProduct(service)

	router.GET("/product/get", controller.GetAll())
	router.POST("/product/add", controller.Store())

	router.Run()
}