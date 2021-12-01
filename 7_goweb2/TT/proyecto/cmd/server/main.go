/*
Se debe separar la estructura del proyecto, como segundo paso se debe generar el paquete
server donde se agregaran las funcionalidades del proyecto que dependan de paquetes
externos y el main del programa.

Dentro del paquete deben estar:
1. El main del programa.
	a. Se debe importar e inyectar el repositorio, servicio y handler
	b. Se debe implementar el router para los diferentes endpoints

2. El paquete handler con el controlador de la entidad seleccionada.
	a. Se debe generar la estructura request
	b. Se debe generar la estructura del controlador que tenga como campo el
	servicio
	c. Se debe generar la función que retorne el controlador
	d. Se deben generar todos los métodos correspondientes a los endpoints
*/

package main

import (
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repository := productos.NewRepository()
	service := productos.NewService(repository)
	controller := handler.NewProducto(service)

	router.GET("/productos", controller.GetAll())
	router.POST("/productos", controller.Store())

	router.Run("localhost:8080")

}
