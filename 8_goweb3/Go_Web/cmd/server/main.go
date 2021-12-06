package main

import (
	"log"
	"os"

	handler "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/docs"
	usuarios "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/internal/usuarios"
	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/pkg/store"
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
		log.Fatal("error al intentar cargar el archivo .env ", err)
	}
	// fmt.Println("----------------  ", os.Getenv("HOST"))

	router := gin.Default()

	db := store.New(store.FileType, "./usuariosSalida.json")
	repository := usuarios.NewRepository(db)
	service := usuarios.NewService(repository)
	controller := handler.NewUsuario(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(TokenAuthMiddleware())

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("usuarios/update", controller.Update())
	router.DELETE("usuarios/delete/:id", controller.Delete())
	router.PATCH("usuarios/patch/:id", controller.EditarNombreEdad())

	router.Run()

}
