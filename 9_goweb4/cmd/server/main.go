package main

import (
	"os"

	"github.com/extmatperez/meli_bootcamp2/9_goweb4/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/docs"
	transaction "github.com/extmatperez/meli_bootcamp2/9_goweb4/internal/transaction"
	store "github.com/extmatperez/meli_bootcamp2/9_goweb4/pkg/store"
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
	transactionURL.POST("/", controller.Store())
	//transactionURL.POST("/", controller.CreateTransaction())
	transactionURL.PUT("/:id", controller.UpdateTransaction())
	transactionURL.PATCH("/:id/:amount", controller.UpdateAmount())
	transactionURL.DELETE("/:id", controller.DeleteTransaction())

	router.Run()
}
