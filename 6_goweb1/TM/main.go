package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int64   `json:"id" binding:"required,min=1,max=16"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

var products []Product

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

func AddPersona(ctx *gin.Context) {
	var pro Product
	err := ctx.ShouldBindJSON(&pro)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})

	} else {

		/*if len(products) == 0 {
			pro.ID = 1
		} else {
			pro.ID = products[len(products)-1].ID + 1
		}*/
		//products = append(products, pro)
		//ctx.JSON(200, pro)
	}
}

func main() {
	router := gin.Default()

	//ruta que devuelve un listado de productos
	router.GET("products", getAll)
	router.POST("add", AddPersona)
	router.Run()
}
