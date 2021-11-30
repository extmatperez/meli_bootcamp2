package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre" binding:"required"`
	Color           string  `json:"color" binding:"required"`
	Precio          float64 `json:"precio" binding:"required"`
	Stock           int     `json:"stock" binding:"required"`
	Codigo          string  `json:"codigo" binding:"required"`
	Publicado       bool    `json:"publicado" binding:"required"`
	FechaDeCreacion string  `json:"fecha_de_creacion" binding:"required"`
}

var products []Product

func LoadProducts(c *gin.Context) {
	data, err := os.ReadFile("../6_goweb1/products.json")
	if err != nil {
		c.String(400, "No se pudo abrir el archivo")
	} else {
		err2 := json.Unmarshal(data, &products)
		if err2 != nil {
			c.String(400, "Problema al convertir el archivo")
		} else {
			c.String(200, "Productos cargados")
		}
	}
}

func GetAll(c *gin.Context) {
	if len(products) == 0 {
		c.String(200, "No hay productos cargados")
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func AddProduct(c *gin.Context) {
	var prod Product
	token := c.GetHeader("token")
	if token != "" {
		if token == "12345" {
			err := c.ShouldBindJSON(&prod)
			if err != nil {
				c.String(400, err.Error())
			} else {
				if len(products) == 0 {
					prod.ID = 1
				} else {
					prod.ID = len(products) + 1 //si existe lo pisa
				}
				products = append(products, prod)
				c.JSON(200, prod)
			}
		} else {
			c.String(401, "Token erroneo")
		}
	} else {
		c.String(401, "no tiene permisos para realizar la peticion solicidata")
	}

}
func main() {
	router := gin.Default()
	products := router.Group("/products")
	products.GET("/", LoadProducts)
	products.GET("/all", GetAll)
	products.POST("/add", AddProduct)
	router.Run(":8080")
}
