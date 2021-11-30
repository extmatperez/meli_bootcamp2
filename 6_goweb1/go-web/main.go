package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

func readData() []Products {

	var list []Products
	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal([]byte(readProducts), &list); err != nil {
		log.Fatal(err)
	}
	return list
}

func getAll(c *gin.Context) {

	var prodList = readData()

	c.JSON(200, gin.H{
		"data": prodList,
	})
}

func filterProducts(ctx *gin.Context) {
	var filtered []*Products
	prodList := readData()

	for i, v := range prodList {
		if ctx.Query("filter") == strconv.FormatBool(v.Published) {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Name {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Color {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Price {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Stock {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Code {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.CreationDate {
			filtered = append(filtered, &prodList[i])
		}
	}

	if len(filtered) != 0 {
		ctx.JSON(200, filtered)
	} else {
		ctx.String(400, "No results found")
	}

}
func main() {
	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Franco!",
		})
	})
	//router.GET("/find/products", filterProducts)
	router.GET("/find/products", filterProducts)
	router.GET("/products", getAll)
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()

}
