/* Ejercicio 1 - Generar paquete internal
Se debe separar la estructura del proyecto y como primer paso generando el paquete internal, en el paquete internal deben estar todas las funcionalidades
que no dependan de paquetes externos.

* Dentro del paquete deben estar las capas:
Servicio, debe contener la lógica de nuestra aplicación.
Se debe crear el archivo service.go.
Se debe generar la interface Service con todos sus métodos.
Se debe generar la estructura service que contenga el repositorio.
Se debe generar una función que devuelva el Servicio.
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
* Repositorio, debe tener el acceso a la variable guardada en memoria.
Se debe crear el archivo repository.go
Se debe crear la estructura de la entidad
Se deben crear las variables globales donde guardar las entidades
Se debe generar la interface Repository con todos sus métodos
Se debe generar la estructura repository
Se debe generar una función que devuelva el Repositorio
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..) */

/*Ejercicio 2 - Generar paquete server

Se debe separar la estructura del proyecto, como segundo paso se debe generar el paquete server donde se agregaran las funcionalidades del proyecto que dependan de paquetes externos y el main del programa.

Dentro del paquete deben estar:
El main del programa.
Se debe importar e inyectar el repositorio, servicio y handler
Se debe implementar el router para los diferentes endpoints
El paquete handler con el controlador de la entidad seleccionada.
Se debe generar la estructura request
Se debe generar la estructura del controlador que tenga como campo el servicio
Se debe generar la función que retorne el controlador
Se deben generar todos los métodos correspondientes a los endpoints
*/

package main

/* import (
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/7_goweb2/afternoon_activities/Exercise_1_&_2/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/7_goweb2/afternoon_activities/Exercise_1_&_2/internal/users"
	"github.com/gin-gonic/gin"
) */

// Creo la función main, agrego mi router y lo inicializo, creo las rutas necesarias y agrego los handlers
func main() {
	/* router := gin.Default()

	repo := users.New_repository()
	service := users.New_service(repo)
	controller := handler.New_user(service)

	router.GET("/users", controller.Get_users())
	router.POST("/users", controller.Post_users())

	router.Run() */
}
