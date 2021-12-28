package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/cmd/server/docs"
	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/cmd/server/handler"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/internal/transaction"
	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("Please set token environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "Token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid token")
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

	router := gin.Default()

	db := store.New(store.FileType, "./transactions.json")

	// router.GET("/hola/:nombre", Greetings)
	// router.GET("/transactions", GetAll)
	// router.GET("/filtrar", FilterQuery)
	// router.GET("/transaction/:id", GetOne)

	// router.POST("/agregarEntidad", addTransaction)

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	controller := handler.NewController(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(TokenAuthMiddleware())

	trans := router.Group("/transactions")
	{
		trans.GET("/get", controller.GetAll())
		trans.POST("/add", controller.Store())
		router.PUT("/transactions/:id", controller.Update())
		router.PATCH("/transactions/:id", controller.UpdateReceptor())
		router.DELETE("/transactions/:id", controller.Delete())
	}

	router.Run()

}
