/* Ejercicio 1 - Mejorar el código
Se debe utilizar Golangci-lint para detectar y corregir todos los posibles errores en el proyecto Go-Web. Para eso se deben seguir los siguientes pasos:
Instalar Golangci-lint.
Ejecutar Golangci-lint (con la configuración por defecto es suficiente).
Si existe algún error, se debe corregir.
*/

package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/12_testing3/morning_activities/Exercise_1_2/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/12_testing3/morning_activities/Exercise_1_2/docs"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/12_testing3/morning_activities/Exercise_1_2/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/12_testing3/morning_activities/Exercise_1_2/pkg/store"
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

	err = router.Run()

	if err != nil {
		log.Fatal(err)
	}
}
