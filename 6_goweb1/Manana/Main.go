package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaccion []struct {
	ID       int    `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
}

type Structure struct {
	Transaccion []interface{}
}

func getAllLindo(c *gin.Context) {
	data, err := os.ReadFile("./Transacciones.json")
	if err != nil {
		c.JSON(400, gin.H{
			"mensaje": "Error en el archivo",
		})
	} else {
		data := string(data)
		decoded := &Structure{}
		//var transacciones []Transaccion
		if err := json.Unmarshal([]byte(data), decoded); err != nil {
			panic(err)
		}
		fmt.Println(decoded)
	}
}

func saludar(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"Saludo":    "Hola Pato",
		"Despedida": "Chau Pato",
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

func main() {

	router := gin.Default()

	router.GET("/Multiple", saludar)
	router.GET("/transaccionesFeo", getAllFeo)
	router.GET("/transaccionesLindo", getAllLindo)

	router.Run()

}
