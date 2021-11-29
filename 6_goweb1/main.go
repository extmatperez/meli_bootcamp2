package main

import (
	"encoding/json"
	"os"
	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion int     `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

func GetAll(c *gin.Context){
	var lista []Transaccion
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &lista)
	c.JSON(200, lista)
}

func main() {

	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, "¡Hola Juampi!")
	})

	router.GET("/transacciones", GetAll)


	router.Run()
	
}