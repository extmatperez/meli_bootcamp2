package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lean1097/meli_bootcamp2/7_goweb2/TM/handlers"
)

func main() {
	fmt.Println("")

	router := gin.Default()
	router.POST("/", handlers.CreateNewProduct)

	router.Run(":3001")
}
