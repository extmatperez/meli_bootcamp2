package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id                int     `json:"id"`
	Nombre            string  `json:"nombre"`
	Color             string  `json:"color"`
	Precio            float64 `json:"precio"`
	Stock             int     `json:"stock"`
	Codigo            string  `json:"codigo"`
	Publicado         bool    `json:"publicado"`
	Fecha_de_creacion string  `json:"fecha_de_creacion"`
}

func saludar(c *gin.Context) {
	datos, _ := os.ReadFile("./productos.json")
	var lista []Productos
	json.Unmarshal(datos, &lista)

	// fmt.Println(lista[0].Nombre)

	c.JSON(http.StatusOK, gin.H{
		"productos": lista,
	})
}
func getAll(c *gin.Context) {

	datos, _ := os.ReadFile("./productos.json")
	var lista []Productos
	json.Unmarshal(datos, &lista)

	c.JSON(http.StatusOK, gin.H{
		"productos": lista,
	})
}

func getById(c *gin.Context) {
	var lista []Productos
	var products []Productos
	datos, err := os.ReadFile("./productos.json")
	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(datos, &lista)
		id, _ := strconv.Atoi(c.Param("id"))
		for _, p := range lista {
			if p.Id == id {
				products = append(products, p)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": products,
		})
	}
}

func getByPublicado(c *gin.Context) {
	var products []Productos
	var filtrados []Productos
	datos, err := os.ReadFile("./productos.json")
	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(datos, &products)
		b, _ := strconv.ParseBool(c.Query("publicado"))
		fmt.Println(b)
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

	router := gin.New()

	router.GET("/saludo", saludar)

	router.GET("/productos", getAll)

	router.GET("/producto/:id", getById)

	router.GET("/publicados", getByPublicado)

	router.Run()
}
