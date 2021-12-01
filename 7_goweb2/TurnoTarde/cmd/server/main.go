package main

import (

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/7_goweb2/TurnoTarde/cmd/server/handler"
	tran "github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/7_goweb2/TurnoTarde/internal/transaccion"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	repo := tran.NewRepository()
	service := tran.NewService(repo)
	controller := handler.NewTransaction(service)
	transaction := server.Group("/transactions")
	{
		transaction.GET("/", controller.GetAll())
		transaction.POST("/", controller.Store())
		transaction.PUT("/:id", controller.Update())
		transaction.PATCH("/:id", controller.UpdateCodigoAndMonto())
		transaction.DELETE("/:id", controller.Delete())
	}	
	server.Run()
}