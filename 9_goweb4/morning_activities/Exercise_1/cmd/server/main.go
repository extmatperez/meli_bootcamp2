/* Ejercicio 1 - Manejo de respuestas genéricas
Se requiere implementar un manejo de respuestas genéricas para enviar siempre el mismo formato en las peticiones. Para lograrlo se deben realizar los
siguientes pasos:
Generar el paquete web dentro del directorio pkg.
Realizar la estructura Response con los capos: code, data y error:
code tendra el codigo de retorno.
data tendrá la estructura que envía la aplicación (en caso que no haya error).
error tendrá el error recibido en formato texto (en caso que haya error).
Desarrollar una función que reciba el code cómo entero, data como interfaz y error como string.
La función debe retornar en base al código, si es una respuesta con el data o con el error.
Implementar esta función en todos los retornos de los controladores, antes de enviar la respuesta al cliente la función debe generar la estructura que definimos.
*/

package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/morning_activities/Exercise_1/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/morning_activities/Exercise_1/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/morning_activities/Exercise_1/pkg/store"
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
