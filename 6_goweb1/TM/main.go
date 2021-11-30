package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        int     `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

func getAll(c *gin.Context) {
	products := []Product{}
	data, err := os.ReadFile("./products.json")
	if err == nil {
		json.Unmarshal(data, &products)
	}
	c.JSON(200, products)
}

func getProduct() []Product {
	products := []Product{}
	data, _ := os.ReadFile("./products.json")
	err := json.Unmarshal(data, &products)

	if err != nil {
		return nil
	} else {
		return products
	}
}

func getById(c *gin.Context) {
	var products []Product = getProduct()
	var productsFiend *Product

	if products != nil {
		idStr := c.Param("ID")
		id, _ := strconv.ParseInt(idStr, 10, 64)
		for _, product := range products {
			if product.ID == id {
				productsFiend = &product
			}
		}
		if productsFiend == nil {
			c.JSON(http.StatusNotFound, "Transaction "+idStr+" not found")
		} else {
			c.JSON(http.StatusOK, productsFiend)
		}
	} else {
		c.JSON(500, nil)
	}
}

func main() {
	router := gin.Default()

	//ruta que devuelve un listado de productos
	router.GET("products", getAll)

	router.Run()
}
