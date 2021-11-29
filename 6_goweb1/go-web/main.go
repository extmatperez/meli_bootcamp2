package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/saludar/:name", saludar)
	router.GET("/saludar", saludar)
	router.GET("/transacciones", getAll)

	router.Run()

}

type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaCreacion     string  `json:"fecha_creacion"`
}

func saludar(c *gin.Context) {
	//queryName := c.Request.URL.Query()  // esto devuelve un map de string string
	queryName := c.Query("name")
	paramName := c.Param("name")

	if queryName != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola " + queryName,
		})
	} else if paramName != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola " + paramName,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola, ingresa tu nombre como query o param",
		})
	}

}

func getAll(c *gin.Context) {

}
