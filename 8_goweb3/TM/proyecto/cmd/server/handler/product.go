/*
Se solicita implementar una funcionalidad que modifique completamente una entidad. Para
lograrlo, es necesario, seguir los siguientes pasos:
	1. Generar un método PUT para modificar la entidad completa
	2. Desde el Path enviar el ID de la entidad que se modificará
	3. En caso de no existir, retornar un error 404
	4. Realizar todas las validaciones (todos los campos son requeridos)
*/

/*
Es necesario implementar una funcionalidad para eliminar una entidad. Para lograrlo, es
necesario, seguir los siguientes pasos:
	1. Generar un método DELETE para eliminar la entidad en base al ID
	2. En caso de no existir, retornar un error 404
*/

/*
Se requiere implementar una funcionalidad que modifique la entidad parcialmente, solo se
deben modificar 2 campos:
	- Si se seleccionó Productos, los campos nombre y precio.
	- Si se seleccionó Usuarios, los campos apellido y edad.
	- Si se seleccionó Transacciones, los campos código de transacción y monto.

Para lograrlo, es necesario, seguir los siguientes pasos:
	1. Generar un método PATCH para modificar la entidad parcialmente, modificando solo 2
	campo (a elección)
	2. Desde el Path enviar el ID de la entidad que se modificara
	3. En caso de no existir, retornar un error 404
	4. Realizar las validaciones de los 2 campos a enviar
*/

package handler

import (
	"strconv"

	productos "github.com/extmatperez/meli_bootcamp2/8_goweb3/TM/proyecto/internal/productos"
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

			response, err := controller.service.Store(nuevoProducto.Nombre, nuevoProducto.Color, nuevoProducto.Precio, nuevoProducto.Stock, nuevoProducto.Codigo, nuevoProducto.Publicado, nuevoProducto.FechaCreacion)

			if err != nil {
				c.String(400, "Hubo un error al guardar los datos: %v", err.Error())
			} else {
				c.JSON(200, response)
			}
		}

	}
}

func (controller *Producto) Update() gin.HandlerFunc {

	return func(c *gin.Context) {

		var productoActualizado request

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "No se pudo leer el id por parámetro")
		} else {

			err = c.ShouldBindJSON(&productoActualizado)

			if err != nil {
				c.String(400, "No se pudo leer el body")
			} else {

				productoActualizado, err := controller.service.Update(id, productoActualizado.Nombre, productoActualizado.Color, productoActualizado.Precio, productoActualizado.Stock, productoActualizado.Codigo, productoActualizado.Publicado, productoActualizado.FechaCreacion)

				if err != nil {
					c.String(404, err.Error())
				} else {
					c.JSON(200, productoActualizado)
				}

			}
		}
	}
}

func (controller *Producto) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "No se pudo leer el id por parámetro")
		} else {

			err := controller.service.Delete(id)

			if err != nil {
				c.String(404, err.Error())
			} else {
				c.String(200, "El producto %v fue eliminado", id)
			}

		}
	}
}

func (controller *Producto) UpdateNombrePrecio() gin.HandlerFunc {

	return func(c *gin.Context) {

		var productoActualizado request

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "No se pudo leer el id por parámetro")
		} else {

			err = c.ShouldBindJSON(&productoActualizado)

			if err != nil {
				c.String(400, "No se pudo leer el body")
			} else {

				if productoActualizado.Nombre == "" || productoActualizado.Precio == "" {
					c.String(404, "El nombre y el precio no pueden estar vacíos")
				} else {
					productoActualizado, err := controller.service.UpdateNombrePrecio(id, productoActualizado.Nombre, productoActualizado.Precio)

					if err != nil {
						c.String(404, err.Error())
					} else {
						c.JSON(200, productoActualizado)
					}
				}

			}
		}
	}
}
