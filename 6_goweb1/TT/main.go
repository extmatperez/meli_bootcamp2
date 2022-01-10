package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float32 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creation_date"`
}

func GetAllProducts(c *gin.Context) {
	prods, _ := os.ReadFile("./products.json")
	var products []Product
	unm := json.Unmarshal(prods, &products)
	if unm != nil {
		c.JSON(500, "Error.")
	} else {
		c.JSON(200, products)
	}
}

func SearchByQuery(c *gin.Context) {
	prods, _ := os.ReadFile("../TM/products.json")
	var todos []Product
	unm := json.Unmarshal(prods, &todos)

	if unm != nil {
		c.JSON(500, "Error.")
	}

	var filtrados []Product

	for i, v := range todos {
		allFilters := true
		if c.Query("name") != "" {
			if c.Query("name") != v.Name {
				allFilters = false
			}
		}
		if c.Query("color") != "" {
			if c.Query("color") != v.Color {
				allFilters = false
			}
		}
		if c.Query("code") != "" {
			if c.Query("code") != v.Code {
				allFilters = false
			}
		}
		if allFilters {
			filtrados = append(filtrados, todos[i])
		}
	}

	if len(filtrados) > 0 {
		c.JSON(200, gin.H{
			"products": filtrados,
		})
	} else {
		c.String(404, "no se encontro nada")
	}
}

func GetProductByID(c *gin.Context) {
	prods, _ := os.ReadFile("../TM/products.json")
	var todos []Product
	unm := json.Unmarshal(prods, &todos)

	if unm != nil {
		c.JSON(500, "Error.")
	}

	id, _ := strconv.Atoi(c.Param("id"))

	for i, v := range todos {
		if id == v.Id {
			c.JSON(200, todos[i])
			break
		}
	}
}

func main() {
	router := gin.Default()
	products := router.Group("/products")
	{
		products.GET("/", GetAllProducts)
		products.GET("/search", SearchByQuery)
		products.GET("/search/:id", GetProductByID)
	}
	router.Run()
}
