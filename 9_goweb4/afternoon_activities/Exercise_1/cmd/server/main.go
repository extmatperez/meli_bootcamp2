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
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/afternoon_activities/Exercise_1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/afternoon_activities/Exercise_1/docs"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/afternoon_activities/Exercise_1/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/9_goweb4/afternoon_activities/Exercise_1/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token is required")
			return
		}
		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// Creo la función main, agrego mi router y lo inicializo, creo las rutas necesarias y agrego los handlers (controllers)
func main() {
	// Levanto las variables de entorno del archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Something went wrong with the .env file")
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	// Creamos una variable que guarde lo que vamos a escribir en nuestro json y se lo pasamos el repo
	db := store.New(store.File_type, "./users.json")
	repo := users.New_repository(db)
	service := users.New_service(repo)
	controller := handler.New_user(service)

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Utilizo el middleware que va a autenticar el token, debajo de las rutas de documentación para que no cree conflicto
	router.Use(TokenAuthMiddleware())

	router.GET("/users", controller.Get_users())
	router.POST("/users", controller.Post_users())
	router.PUT("/users/:id", controller.Update_users())
	router.PATCH("/users/:id", controller.Update_users_first_name())
	router.DELETE("/users/:id", controller.Delete_users())

	router.Run()
}
