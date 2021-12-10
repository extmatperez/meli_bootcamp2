package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID      int    `json:"Id"`
	Name    string `json:"Name"`
	Color   string `json:"Color"`
	Price   string `json:"Price"`
	Stock   int    `json:"Stock"`
	Code    string `json:"Code"`
	Publish bool   `json:"Publish"`
	Date    string `json:"Date"`
}

func GetAll(c *gin.Context) {
	manejador, _ := ioutil.ReadFile("Products.json")

	var p []Product

	if err2 := json.Unmarshal(manejador, &p); err2 != nil {
		log.Fatal(err2)
	}
	c.JSON(200, p)
}

func AgregarProducto(c *gin.Context) {

}

func main() {
	router := gin.Default()

	router.GET("Hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Andres",
		})
	})

	router.GET("Productos", GetAll)

	router.POST("/productos", AgregarProducto)

	router.Run()
}
