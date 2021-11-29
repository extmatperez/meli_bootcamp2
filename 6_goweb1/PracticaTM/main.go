package main

import (
	"encoding/json"
	"log"
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

func buscarTransaccion(c *gin.Context) {
	var transac Transaccion

	if c.BindJSON(&transac) == nil {
		log.Println("Bind por JSON")
		log.Println("ID de transaccion: ", transac.Id)
		log.Println("Codigo de transaccion: ", transac.CodTransaccion)
		c.String(http.StatusOK, "(Query JSON) - Transaccion: %s, ID: %s\n", transac.CodTransaccion, transac.Id)
	} else {
		c.String(404, "La transaccion no existe")
	}
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

	router.GET("/transacciones/:id", buscarTransaccion)

	router.Run()
}
