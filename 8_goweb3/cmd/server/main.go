package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/docs"
	tran "github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/internal/transaccion"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle Transactions
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main(){
	server := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}


	//inicializaciones
	db := store.New(store.FileType,os.Getenv("STOREPATH"))
	repo := tran.NewRepository(db)
	service := tran.NewService(repo)
	controller := handler.NewTransaction(service)

	
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   



	transaction := server.Group("/transactions")



	{
		//get
		transaction.GET("/", controller.GetAll())
		transaction.GET("/:id", controller.GetTransactionById())
		transaction.GET("/filtros", controller.GetTransactionsExlusive())

		//post
		transaction.POST("/", controller.Store())


		//put
		transaction.PUT("/:id", controller.Update())

		//patch
		transaction.PATCH("/:id", controller.UpdateCodigoAndMonto())


		//delete
		transaction.DELETE("/:id", controller.Delete())
	}	
	server.Run()
}