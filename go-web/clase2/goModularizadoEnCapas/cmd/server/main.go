package main

import (
	productos "github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/go-web/clase2/goModularizadoEnCapas/Internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/go-web/clase2/goModularizadoEnCapas/cmd/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := productos.NewRepository()
	service := productos.NewService(repo)
	controller := handler.NewPersona(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/modificar/:id", controller.Update())
	//router.PATCH("/modificarNombre/:id", controller.UpdateName())
	//router.DELETE("/delete/:id", controller.delete())
	router.Run()
}
