package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       bool    `json:"stock"`
	Code        string  `json:"code"`
	Published   bool    `json:"published"`
	DateCreated string  `json:"date_created"`
}

func getAll(c *gin.Context) {
	//Approach 1
	p1 := Product{10, "Automovil", "Blanco", 150000.98, true, "as4asfd", true, "29/11/2021"}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	} 
	c.Data(http.StatusOK, gin.MIMEJSON, jsonData)

	/*
	//Approach 2
	datos, _ := os.ReadFile("./products.json")
	var lista []Product
	json.Unmarshal(datos, &lista)

	c.JSON(http.StatusOK, gin.H{
		"productos": lista,
	})
	*/
}

func main() {

	router := gin.Default()

	//grouping
	productsEP := router.Group("/products")
	{
		productsEP.GET("/GetAll", getAll)
	}

	router.Run()

}