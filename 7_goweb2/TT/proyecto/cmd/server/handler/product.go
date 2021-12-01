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

package handler

import (
	productos "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

type Producto struct {
	service productos.Service
}

func NewProducto(s productos.Service) *Producto {
	return &Producto{service: s}
}

func (controller *Producto) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		productos, err := controller.service.GetAll()

		if err != nil {
			c.String(400, "Hubo un error: %v", err.Error())
		} else {
			c.JSON(200, productos)
		}

	}
}

func (controller *Producto) Store() gin.HandlerFunc {

	return func(c *gin.Context) {

		var nuevoProducto request

		err := c.ShouldBindJSON(&nuevoProducto)

		if err != nil {
			c.String(400, "Hubo un error al recibir los datos: %v", err.Error())
		} else {

			response, err := controller.service.Store(nuevoProducto.Nombre, nuevoProducto.Color, nuevoProducto.Precio, nuevoProducto.Stock, nuevoProducto.Codigo, nuevoProducto.Publicado, nuevoProducto.Nombre)

			if err != nil {
				c.String(400, "Hubo un error al guardar los datos: %v", err.Error())
			} else {
				c.JSON(200, response)
			}
		}

	}
}
