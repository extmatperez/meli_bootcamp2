/* Ejercicio 2 - Guardar información
Se debe implementar la funcionalidad para guardar la información de la petición en un
archivo json, para eso se deben realizar los siguientes pasos:
1. En lugar de guardar los valores de nuestra entidad en memoria, se debe crear un
archivo; los valores que se vayan agregando se guardan en él. */

package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/8_goweb3/afternoon_activities/Exercise_2/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/8_goweb3/afternoon_activities/Exercise_2/internal/users"
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
