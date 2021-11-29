package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lean1097/meli_bootcamp2/6_goweb1/TT/handlers"
)

func main() {
	fmt.Println("New server")

	router := gin.Default()

	router.GET("/", handlers.InitHandler)
	router.GET("/products", handlers.GetAllProducts)
	router.GET("/product/:id", handlers.GetProductByID)
	router.GET("/transactions", handlers.GetAllTransactions)
	router.GET("/users", handlers.GetAllUsers)

	router.Run(":3000")
}
