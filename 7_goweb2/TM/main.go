package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

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

var products []Product

func GetProducts(c *gin.Context) {
	if len(products) > 0 {
		c.JSON(200, products)
	} else {
		c.JSON(200, "no se encontro nada")
	}
}

func filtrar(sliceProducts []Product, campo string, valor string) []Product {
	var filtrado []Product

	var prod Product
	tipos := reflect.TypeOf(prod)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		if strings.ToLower(tipos.Field(i).Name) == campo {
			break
		}
	}

	for _, v := range sliceProducts {
		cadena := fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor) {
			filtrado = append(filtrado, v)
		}
	}

	return filtrado
}

func FilterProducts(ctx *gin.Context) {
	var etiquetas []string
	etiquetas = append(etiquetas, "name", "color", "code", "creation_date")

	var productosFiltrados []Product

	productosFiltrados = products

	for _, v := range etiquetas {
		if len(ctx.Query(v)) != 0 && len(productosFiltrados) != 0 {
			productosFiltrados = filtrar(productosFiltrados, v, ctx.Query(v))
		}
	}

	if len(productosFiltrados) == 0 {
		ctx.String(200, "No hay coincidencias")
	} else {
		ctx.JSON(200, productosFiltrados)
	}
}

func AddProduct(c *gin.Context) {
	var prod Product

	err := c.ShouldBindJSON(&prod)
	if err != nil {
		c.JSON(400, err.Error())
	} else {
		prod.Id = len(products) + 1
		products = append(products, prod)
		c.JSON(200, prod)
	}
}

func AddJson(c *gin.Context) {
	data, err := os.ReadFile("products.json")
	if err != nil {
		c.String(200, "error con el json")
	} else {
		json.Unmarshal(data, &products)
		c.JSON(200, products)
	}

}

func main() {
	router := gin.Default()
	products := router.Group("/products")
	{
		products.GET("/", GetProducts)
		products.GET("/filtros", FilterProducts)
		products.POST("/add", AddProduct)
		products.POST("/addjson", AddJson)
	}
	router.Run()
}
