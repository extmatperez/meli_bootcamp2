package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/11_testing2/PracticaTM/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/11_testing2/PracticaTM/docs"
	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/11_testing2/PracticaTM/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/11_testing2/PracticaTM/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//Estas funciones son para el middleware. A la api le tenemos que indicar que use este middleware para autenticar. Si está todo bien
// le permite seguir con la ejecución, sino le da un error y no deja seguir a la api
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")
	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("api_token")
		if token == "" {
			respondWithError(c, 401, "API token required")
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
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// A partir de aca, con esta funcion Use seteamos el middleware para todos los endpoints de abajo.
	// De ahora en más, para cualquier endpoint que utilice "router" se va a autenticar con el Middleware.
	// Ya no hay que definir la verificacion del token en cada endpoint.
	router.Use(TokenAuthMiddleware())

	db := store.New(store.FileType, "./transaccionesSalida.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	transac := handler.NewTransaccion(service)

	groupTransac := router.Group("/transactions")

	//Endpoints GET
	groupTransac.GET("/", transac.GetAll())

	groupTransac.GET("/:id", transac.Search())
	groupTransac.GET("/filter", transac.Filter())

	// //Endpoints POST
	groupTransac.POST("/load", transac.Store())

	// Endpoints Put
	groupTransac.PUT("/:id", transac.Update())

	// Endpoint delete
	groupTransac.DELETE("/:id", transac.Delete())

	// Endpoint Patch
	groupTransac.PATCH("/:id", transac.UpdateCodigoYMonto())

	router.Run()
}
