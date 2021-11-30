package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

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

var productos []Product

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
func validaVacio(p Product) {

}

func createProduct(c *gin.Context) {
	var pro Product
	var etiquetas []string
	etiquetas = append(etiquetas, "id", "name", "color", "price", "stock", "code", "Publish", "date")
	err := c.ShouldBindJSON(&pro)
	var faltantes []string

	for i := range etiquetas {
		var cadena string
		cadena = fmt.Sprintf("%v", reflect.ValueOf(pro).Field(i).Interface())
		println(cadena)
		if cadena == "" || cadena == "0" {
			println(cadena)
			cadena2 := fmt.Sprintf("%v", etiquetas[i])
			//c.String(401, "Falta el campo :"+cadena2)
			faltantes = append(faltantes, cadena2)

		}
	}
	if len(faltantes) > 0 {
		c.String(401, "Faltan los campos :"+fmt.Sprintf("%v", faltantes[1:]))
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"Error ": err.Error(),
		})
	} else {
		if len(productos) == 0 {
			pro.Id = 1
		} else {
			pro.Id = productos[len(productos)-1].Id + 1
		}
		productos = append(productos, pro)
		c.JSON(200, pro)
	}
}

func main() {

	s := gin.Default()

	s.GET("/saludo/:name", saludar)
	s.GET("/publicados", FiltrarProductos)
	s.GET("/productos/:id", BuscarProducto)
	//
	s.POST("/crear", createProduct)

	s.Run()

}
