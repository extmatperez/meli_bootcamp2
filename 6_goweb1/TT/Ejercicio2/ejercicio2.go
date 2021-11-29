package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Stock      int    `json:"stock"`
	Code       string `json:"code"`
	Published  bool   `json:"published"`
	Created_at string `json:"created_at"`
}

func getAllProducts(ctx *gin.Context) {
	bytes, err := os.ReadFile("../../TM/Ejercicio1/products.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "No se pudieron cargar los productos",
		})
		return
	}

	var allProducts []Product
	errUnmarshal := json.Unmarshal(bytes, &allProducts)

	if errUnmarshal != nil {
		ctx.JSON(500, gin.H{
			"error": "Error parseando el JSON de productos",
		})
		return
	}

	nameQuery := ctx.Query("name")
	colorQuery := ctx.Query("color")
	stockQuery := ctx.Query("stock")
	codeQuery := ctx.Query("code")
	publishedQuery := ctx.Query("published")
	createdAtQuery := ctx.Query("created_at")

	var products []Product

	for i := 0; i < len(allProducts); i++ {
		includesAllFilters := true
		if nameQuery != "" {
			if !strings.EqualFold(allProducts[i].Name, nameQuery) {
				includesAllFilters = false
			}
		}
		if colorQuery != "" {
			if !strings.EqualFold(allProducts[i].Color, colorQuery) {
				includesAllFilters = false
			}
		}
		if stockQuery != "" {
			stock, err := strconv.Atoi(stockQuery)
			if err == nil {
				if allProducts[i].Stock != stock {
					includesAllFilters = false
				}
			}
		}
		if codeQuery != "" {
			if !strings.EqualFold(allProducts[i].Code, codeQuery) {
				includesAllFilters = false
			}
		}
		if publishedQuery != "" {
			if publishedQuery == "true" || publishedQuery == "false" {
				var published bool
				if publishedQuery == "true" {
					published = true
				} else if publishedQuery == "false" {
					published = false
				}

				if allProducts[i].Published != published {
					includesAllFilters = false
				}
			}
		}
		if createdAtQuery != "" {
			if !strings.EqualFold(allProducts[i].Created_at, createdAtQuery) {
				includesAllFilters = false
			}
		}

		if includesAllFilters {
			products = append(products, allProducts[i])
		}
	}

	ctx.JSON(200, gin.H{
		"products": products,
	})
}

func getProduct(ctx *gin.Context) {
	productId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "ID invalido",
		})
		return
	}

	bytes, err := os.ReadFile("../../TM/Ejercicio1/products.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "No se pudieron cargar los productos",
		})
		return
	}

	var allProducts []Product
	errUnmarshal := json.Unmarshal(bytes, &allProducts)

	if errUnmarshal != nil {
		ctx.JSON(500, gin.H{
			"error": "Error parseando el JSON de productos",
		})
		return
	}

	var product Product

	for i := 0; i < len(allProducts); i++ {
		if allProducts[i].Id == productId {
			product = allProducts[i]
			break
		}
	}

	if product == (Product{}) {
		ctx.JSON(404, gin.H{
			"message": "Producto no encontrado",
		})
		return
	}

	ctx.JSON(200, product)
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hola " + name,
		})
	})

	products := router.Group("/products")
	{
		products.GET("/", getAllProducts)
		products.GET("/:id", getProduct)
	}

	router.Run()
}
