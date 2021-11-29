package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Fecha struct {
	Dia, Mes, Anio int
}
type Transaccion struct {
	Id             int
	CodTransaccion string
	Moneda         string
	Monto          float64
	Emisor         string
	Receptor       string
	FechaTrans     string
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola Facundo",
		})
	})

	router.Run(":8081")
}
