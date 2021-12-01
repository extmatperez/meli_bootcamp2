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
	Id                int     `json:"id" required:"true"`
	Nombre            string  `json:"nombre"`
	Color             string  `json:"color"`
	Precio            float64 `json:"precio"`
	Stock             int     `json:"stock"`
	Codigo            string  `json:"codigo"`
	Publicado         bool    `json:"publicado"`
	Fecha_de_creacion string  `json:"fecha_de_creacion"`
}

var lista []Productos

func saludar(c *gin.Context) {
	datos, _ := os.ReadFile("./productos.json")
	json.Unmarshal(datos, &lista)

	c.JSON(http.StatusOK, gin.H{
		"productos": lista,
	})
}
func getAll(c *gin.Context) {

	datos, _ := os.ReadFile("./productos.json")
	json.Unmarshal(datos, &lista)

	c.JSON(http.StatusOK, gin.H{
		"productos": lista,
	})
}

func getById(c *gin.Context) {
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
		// fmt.Println(b)
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

func addProductos(c *gin.Context) {
	var producto Productos
	err := c.ShouldBindJSON(&producto)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		if len(lista) == 0 {
			producto.Id = 1
		} else {
			producto.Id = lista[len(lista)-1].Id + 1
		}
		lista = append(lista, producto)
		c.JSON(http.StatusOK, lista)
	}
}

func token(c *gin.Context) {
	token := c.GetHeader("token")

	if token != "" && token == "123456" {
		c.JSON(http.StatusOK, gin.H{
			"success": "Acceso autorizado",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Token incorrecto",
		})
		return
	}
}

func main() {

	router := gin.New()

	router.GET("/saludo", saludar)

	router.GET("/productos", getAll)

	router.GET("/producto/:id", getById)

	router.GET("/publicados", getByPublicado)

	router.POST("/agregar", token, addProductos)

	router.Run()
}
