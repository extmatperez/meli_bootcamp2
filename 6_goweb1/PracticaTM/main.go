package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Fecha struct {
	Dia, Mes, Anio int
}
type Transaccion struct {
	Id             int     `json:"id"`
	CodTransaccion string  `json:"cod_transaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         string  `json:"emisor"`
	Receptor       string  `json:"receptor"`
	FechaTrans     string  `json:"fecha_trans"`
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola Facundo",
		})
	})

	data, err := os.ReadFile("./transactions.json")

	if err != nil {
		panic("error abriendo el archivo")
	}
	var transacciones []Transaccion

	err = json.Unmarshal(data, &transacciones)

	if err != nil {
		panic("error haciendo el unmarshal")
	}

	router.GET("/transacciones", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GetAll": transacciones,
		})
	})

	router.Run()
}
