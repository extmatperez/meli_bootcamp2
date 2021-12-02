package main

import (
	handler "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/cmd/server/handler"
	usuarios "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repository := usuarios.NewRepository()
	service := usuarios.NewService(repository)
	controller := handler.NewUsuario(service)

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("usuarios/update", controller.Update())
	router.DELETE("usuarios/delete/:id", controller.Delete())
	router.PATCH("usuarios/patch/:id", controller.EditarNombreEdad())

	router.Run()

}
