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

func getOne(c *gin.Context) {
	parameter := c.Param("id")

	var prodList = readData()

	//var prod []Products
	prods := []Products{}
	//is_Product := false

	for _, v := range prodList {
		if strconv.Itoa(v.ID) == parameter {
			prods = append(prods, v)
			//is_Product = true
		}
	}

	if len(prods) > 0 {
		c.JSON(200, prods)
	} else {
		c.String(400, "No product found")
	}
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

func addProduct(ctx *gin.Context) {
	var prod Products
	prodList := readData()
	err := ctx.ShouldBindJSON(&prod)
	lenthProds := len(prodList)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		if lenthProds == 0 {
			prod.ID = 1
		} else {
			prod.ID = prodList[lenthProds-1].ID + 1
		}
		prodList = append(prodList, prod)
		ctx.JSON(200, prod)
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

	products := router.Group("/products")
	products.GET("/find", filterProducts)
	products.GET("/:id", getOne)
	products.GET("/", getAll)
	products.POST("/", addProduct)
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()

}
