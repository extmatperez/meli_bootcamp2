package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/cmd/server/handler"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/internal/transaction"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.GET("/hola/:nombre", Greetings)
	// router.GET("/transactions", GetAll)
	// router.GET("/filtrar", FilterQuery)
	// router.GET("/transaction/:id", GetOne)

	// router.POST("/agregarEntidad", addTransaction)

	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	controller := handler.NewController(service)

	router.GET("/transactions/get", controller.GetAll())
	router.POST("/transactions/add", controller.Store())
	// router.PUT("/transactions/:id", controller.Update())
	// router.PATCH("/transactions/:id", controller.UpdateNombre())
	// router.DELETE("/transactions/:id", controller.Delete())

	router.Run(":8080")

}
