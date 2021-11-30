package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Color     string `json:"color"`
	Precio    string `json:"precio"`
	Stock     int    `json:"stock"`
	Codigo    string `json:"codigo"`
	Publicado bool   `json:"publicado"`
	Creado    string `json:"creado"`
}

var productos []Producto

func main() {
	router := gin.Default()
	router.GET("/hola", saludar)
	router.GET("/productos", getAll)
	router.POST("/productos", crearProducto)
	router.Run()
}

func crearProducto(c *gin.Context) {
	token := c.GetHeader("token")
	tokenInfo := checkToken(token)
	if tokenInfo == "OK" {
		var req Producto
		err := c.ShouldBindJSON(&req)
		valores := reflect.ValueOf(req)
		for i := 0; i < valores.NumField(); i++ {
			valorDeCampo := valores.Field(i)
			if fmt.Sprintf("%s", valorDeCampo) == "" {
				key := reflect.TypeOf(req).Field(i).Name
				descr := fmt.Sprintf("El campo %s es requerido", key)
				c.JSON(400, descr)
				return
			}
		}

		if err != nil {
			fmt.Println(err)
			c.JSON(400, "Ha ocurrido un error")
			return
		}
		productos = append(productos, req)
		c.JSON(201, productos)
		return
	}
	c.JSON(201, tokenInfo)

}

func checkToken(token string) string {
	if token != "" {
		if token == "123456" {
			return "OK"
		}
		return "Token incorrecto"
	} else {
		return "no se ha ingresado un token"
	}
}

func saludar(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola, Matias",
	})
}

func getAll(c *gin.Context) {
	productos, _ := os.ReadFile("productos.json")
	var filtrados []Producto
	var punteroProductos []Producto
	err := json.Unmarshal(productos, &punteroProductos)
	for _, e := range punteroProductos {
		idQuery, _ := strconv.Atoi(c.Query("id"))
		fmt.Println(c.Query("id"), e.ID, idQuery == e.ID)
		if idQuery == e.ID {
			filtrados = append(filtrados, e)
		}
	}

	if err != nil {
		c.JSON(500, "Ha ocurrido un error")
		return
	}
	c.JSON(200, filtrados)
}
