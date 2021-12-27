package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/docs"
	payments "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/internal/payments"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/17_storage1/TT/EjercicioClase/pkg/store"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	// Toma el token del .env
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Por favor, setear la variable de entorno TOKEN.")
	}

	return func(c *gin.Context) {
		// Toma el token del header en el postman.
		token := c.GetHeader("TOKEN")

		if token == "" {
			respondWithError(c, 401, "API Token requerido.")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "API Token inválido.")
			return
		}

		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el archivo .env")
	}

	db := store.New(store.FileType, "./payments.json")
	repository := payments.NewRepository(db)
	service := payments.NewService(repository)
	controller := handler.NewPayment(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(TokenAuthMiddleware())

	payments := router.Group("/payments")
	{
		payments.GET("/get", controller.GetAll())
		payments.POST("/", controller.Store())
		payments.PUT("/:id", controller.Update())
		payments.PATCH("/code/:id", controller.UpdateCodigo())
		payments.PATCH("/amount/:id", controller.UpdateMonto())
		payments.DELETE("/:id", controller.Delete())
	}

	err0 := router.Run(":8080")
	if err0 != nil {
		log.Fatal("Error al intentar correr el router.")
	}
}
