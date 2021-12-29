package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/docs"
	productos "github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"Error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {

	tokenENV := os.Getenv("TOKEN")

	if tokenENV == "" {
		log.Fatal("Por favor seteá una token en .env")
	}

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "Token no enviado")
			return
		}

		if token != tokenENV {
			respondWithError(c, 401, "Token inválido")
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

	db := store.New(store.FileType, "/Users/beconti/Desktop/meli_bootcamp2/7_goweb2/productos.json")

	repository := productos.NewRepository(db)
	service := productos.NewService(repository)
	controller := handler.NewProducto(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(TokenAuthMiddleware())

	productos := router.Group("/productos")
	{
		productos.GET("/", controller.GetAll())
		productos.POST("/", controller.Store())
		productos.PUT("/:id", controller.Update())
		productos.DELETE("/:id", controller.Delete())
		productos.PATCH("/:id", controller.UpdateNombrePrecio())
	}

	err = router.Run("localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
}
