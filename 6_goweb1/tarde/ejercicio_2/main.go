package main

import (
	"encoding/json"
	"net/http"
	"os"

	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Code    string  `json:"code"`
	Publish bool    `json:"publish"`
	Date    string  `json:"date"`
}

func GetAll(c *gin.Context) {
	archivo, _ := os.ReadFile("./products.json")
	var listProducts []Product
	json.Unmarshal(archivo, &listProducts)

	c.JSON(http.StatusOK, gin.H{
		"products": listProducts,
	})
}

func FiltrarProductos(ctx *gin.Context) {
	archivo, _ := os.ReadFile("./products.json")
	var listProducts []Product
	json.Unmarshal(archivo, &listProducts)
	var filtrados []*Product

	for i, v := range listProducts {
		if ctx.Query("filtro") == strconv.FormatBool(v.Publish) {
			filtrados = append(filtrados, &listProducts[i])
		}
	}

	if len(filtrados) == 0 {
		ctx.String(400, "No se encontró Productos con esa caracteristica")
	} else {
		ctx.JSON(200, filtrados)
	}

}

func BuscarProducto(ctx *gin.Context) {
	archivo, _ := os.ReadFile("./products.json")
	var listProducts []Product
	json.Unmarshal(archivo, &listProducts)

	parametro := ctx.Param("id")
	var emp Product
	se := false
	for _, v := range listProducts {
		if strconv.Itoa(v.Id) == parametro {
			emp = v
			se = true
			break
		}
	}

	if se {
		ctx.JSON(200, emp)
	} else {
		ctx.String(404, "No se encontro el producto %s", parametro)
	}

}
func saludar(c *gin.Context) {
	parametro := c.Param("name")
	c.String(http.StatusOK,
		"Hello "+parametro+" !!",
	)

}

func main() {

	s := gin.Default()
	// s.GET("/saludo", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hola Diego!!",
	// 	})
	// })
	s.GET("/saludo/:name", saludar)
	s.GET("/publicados", FiltrarProductos)
	s.GET("/productos/:id", BuscarProducto)
	s.Run()

}
