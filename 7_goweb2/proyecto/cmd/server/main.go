package main

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := internal.NewRepository()
	service := internal.NewService(repo)
	controller := handler.NewTransaction(service)

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	routerTransactions := router.Group("/transactions")
	{
		routerTransactions.GET("/get", controller.GetAll())
		routerTransactions.GET("/get/:id", controller.GetTransactionByID())
		routerTransactions.POST("/add", controller.Store())
	}
	router.Run()

	// //List Endpoints
	// rutas := router.Routes()
	// for _, r := range rutas {
	// 	fmt.Printf("Method: %s \t\t Path: %s\n", r.Method, r.Path)
	// }

}
