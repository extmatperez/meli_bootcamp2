package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}

func saludar(c *gin.Context) {
	name := c.Param("fulano")
	c.JSON(200, gin.H{
		"message": "Hola " + name,
	})
}
func GetAll(c *gin.Context) {
	var products []Products
	ps, err := os.ReadFile("./6_goweb1/go-web/archivos/products.json")
	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(ps, &products)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": products,
	})
}
func GetById(c *gin.Context) {
	var products []Products
	var product Products
	ps, err := os.ReadFile("./6_goweb1/go-web/archivos/products.json")
	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(ps, &products)
		id, _ := strconv.Atoi(c.Param("id"))
		for _, p := range products {
			if p.ID == id {
				product = p
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": product,
		})
	}
}
func GetByPublicado(c *gin.Context) {
	var products []Products
	var filtrados []Products
	ps, err := os.ReadFile("./6_goweb1/go-web/archivos/products.json")
	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(ps, &products)
		b, _ := strconv.ParseBool(c.Query("publicado"))
		for _, p := range products {
			if b == p.Publicado {
				filtrados = append(filtrados, p)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": filtrados,
		})
	}
}
func main() {

	router := gin.Default()
	router.GET("/hola/:fulano", saludar)
	router.GET("/products", GetAll)
	router.GET("product/:id", GetById)             //localhost:8080/product/1
	router.GET("productpublicado", GetByPublicado) // localhost:8080/productpublicado?publicado=false
	router.Run()
}
