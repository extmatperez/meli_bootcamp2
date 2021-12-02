package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/cmd/server/handler"
	tran "github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/internal/transaccion"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	server := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}


	//inicialicaciones
	db := store.New(store.FileType,os.Getenv("STOREPATH"))
	repo := tran.NewRepository(db)
	service := tran.NewService(repo)
	controller := handler.NewTransaction(service)
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