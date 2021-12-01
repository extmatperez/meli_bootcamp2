package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// readData()
	// length := len(prodList) - 1
	// lastID = prodList[length].ID
	// tokenPrueba = "1234"

	/////////////////////// start server /////////////////////////////
	router := gin.Default()

	products := router.Group("/products")

	products.GET("/", getAll)
	products.GET("/products/:id", getProductbyID)
	// products.GET("/products/filter/select", getbyFilter)

	products.POST("/addproduct", addProduct)

	router.Run()

}
