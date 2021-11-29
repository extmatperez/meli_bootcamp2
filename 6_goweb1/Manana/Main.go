package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	ID       int    `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
}

func saludar(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"Saludo":    "Hola Pato",
		"Despedida": "Chau Pato",
	})
}

func saludar2(c *gin.Context) {
	nombre := c.Param("nombre")
	fmt.Println(nombre)
	c.JSON(http.StatusOK, gin.H{
		"Saludo": "Hola " + nombre,
	})
}

func getAllFeo(c *gin.Context) {
	fmt.Println("Hola a GetAll")
	data, err := os.ReadFile("./Transacciones.json")
	if err != nil {
		c.JSON(400, gin.H{
			"mensaje": "Error en el archivo",
		})
	} else {
		c.JSON(200, gin.H{
			"mensaje": "Todo bien",
			"data":    string(data),
		})
	}
}

func getAllLindo(c *gin.Context) {

	res, err := os.ReadFile("./Transacciones.json")
	if err != nil {
		c.JSON(400, gin.H{
			"mensaje": "Error en el archivo",
		})
	} else {
		str := string(res)
		fmt.Println(str)
		var transferencias []Transaccion
		json.Unmarshal(res, &transferencias)

		var finalText string

		for _, t := range transferencias {
			newText := fmt.Sprintf("\n La transaccion %v por un monto de %v %v la genero %v para %v \n", t.ID, t.Monto, t.Moneda, t.Emisor, t.Receptor)
			finalText = finalText + newText
		}
		fmt.Println(finalText)
		c.JSON(200, finalText)

	}
}

func main() {

	router := gin.Default()

	router.GET("/saludo", saludar)
	router.GET("/saludoLindo/:nombre", saludar2)
	router.GET("/transaccionesFeo", getAllFeo)
	router.GET("/transaccionesLindo", getAllLindo)

	router.Run()

}
