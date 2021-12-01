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

package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TM/proyecto/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/8_goweb3/TM/proyecto/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	repository := productos.NewRepository()
	service := productos.NewService(repository)
	controller := handler.NewProducto(service)

	router.GET("/productos", controller.GetAll())
	router.POST("/productos", controller.Store())
	router.PUT("/productos/:id", controller.Update())
	router.DELETE("/productos/:id", controller.Delete())
	router.PATCH("/productos/:id", controller.UpdateNombrePrecio())

	router.Run("localhost:8080")

}
