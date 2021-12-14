package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_morning/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_morning/docs"
	transaction "github.com/extmatperez/meli_bootcamp2/12_testing3/turn_morning/internal/transaction"
	store "github.com/extmatperez/meli_bootcamp2/12_testing3/turn_morning/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	godotenv.Load()

	router := gin.Default()
	db := store.New(store.FileType, "./transaction.json")
	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	controller := handler.NewTransaction(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transactionURL := router.Group("/transactions")

	transactionURL.GET("/", controller.GetAll())
	transactionURL.GET("/:id", controller.GetByID())
	transactionURL.GET("/receivers/:receiver", controller.GetByReceiver())
	transactionURL.POST("/", TokenAuthMiddleware(), controller.Store())
	//transactionURL.POST("/", controller.CreateTransaction())
	transactionURL.PUT("/:id", TokenAuthMiddleware(), controller.UpdateTransaction())
	transactionURL.PATCH("/:id/:amount", TokenAuthMiddleware(), controller.UpdateAmount())
	transactionURL.DELETE("/:id", TokenAuthMiddleware(), controller.DeleteTransaction())

	router.Run(":9090")
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

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
