/* Ejercicio 1 - Generar método PUT

Se solicita implementar una funcionalidad que modifique completamente una entidad. Para
lograrlo, es necesario, seguir los siguientes pasos:
1. Generar un método PUT para modificar la entidad completa
2. Desde el Path enviar el ID de la entidad que se modificará
3. En caso de no existir, retornar un error 404
4. Realizar todas las validaciones (todos los campos son requeridos) */

package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/8_goweb3/morning_activities/Exercise_1/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/8_goweb3/morning_activities/Exercise_1/internal/users"
	"github.com/gin-gonic/gin"
)

// Creo la función main, agrego mi router y lo inicializo, creo las rutas necesarias y agrego los handlers (controllers)
func main() {
	router := gin.Default()

	repo := users.New_repository()
	service := users.New_service(repo)
	controller := handler.New_user(service)

	router.GET("/users", controller.Get_users())
	router.POST("/users", controller.Post_users())
	router.PUT("/users/:id", controller.Update_users())
	router.PATCH("/users/:id", controller.Update_users_first_name())
	router.DELETE("/users/:id", controller.Delete_users())

	router.Run()
}
