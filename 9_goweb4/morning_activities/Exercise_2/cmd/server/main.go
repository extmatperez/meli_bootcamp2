/* Ejercicio 2 - Documentación
Se requiere implementar la documentación del proyecto utilizando swagger, en el mismo deben estar todos los endpoints documentados. Para lograrlo se debe tener en cuenta:
Utilizar el paquete swaggo.
Generar las anotaciones en la función main.
Realizar las anotaciones en los controladores.
Desarrollar el paquete docs con swaggo en base a las anotaciones.
Agregar un endpoint mediante GET para visualizar la documentación generada.

*/

package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/morning_activities/Exercise_2/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/morning_activities/Exercise_2/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/morning_activities/Exercise_2/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Creo la función main, agrego mi router y lo inicializo, creo las rutas necesarias y agrego los handlers (controllers)
func main() {
	// Levanto las variables de entorno del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something went wrong with the .env file")
	}

	router := gin.Default()

	// Creamos una variable que guarde lo que vamos a escribir en nuestro json y se lo pasamos el repo
	db := store.New(store.File_type, "./users.json")
	repo := users.New_repository(db)
	service := users.New_service(repo)
	controller := handler.New_user(service)

	router.GET("/users", controller.Get_users())
	router.POST("/users", controller.Post_users())
	router.PUT("/users/:id", controller.Update_users())
	router.PATCH("/users/:id", controller.Update_users_first_name())
	router.DELETE("/users/:id", controller.Delete_users())

	router.Run()
}
