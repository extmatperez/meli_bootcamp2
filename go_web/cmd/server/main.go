package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/cmd/server/handler"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/internal/transaction"
	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	trans := router.Group("/transactions")
	trans.GET("/get", controller.GetAll())
	trans.POST("/add", controller.Store())
	router.PUT("/transactions/:id", controller.Update())
	router.PATCH("/transactions/:id", controller.UpdateReceptor())
	router.DELETE("/transactions/:id", controller.Delete())

	router.Run(":8080")

}
