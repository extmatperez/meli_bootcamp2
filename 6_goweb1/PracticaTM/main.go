package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

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

func buscarTransaccion(ctx *gin.Context) {
	var transac Transaccion

	parametro := ctx.Param("id")
	se := false
	for _, valor := range transacciones {
		if strconv.Itoa(valor.Id) == parametro {
			transac = valor
			se = true
			break
		}
	}

	if se {
		ctx.JSON(200, transac)
	} else {
		ctx.String(404, "No se encontro la transacción ""%s""", parametro)
	}

	// if c.BindJSON(&transac) == nil {
	// 	log.Println("Bind por JSON")
	// 	log.Println("ID de transaccion: ", transac.Id)
	// 	log.Println("Codigo de transaccion: ", transac.CodTransaccion)
	// 	c.String(http.StatusOK, "(Query JSON) - Transaccion: %s, ID: %s\n", transac.CodTransaccion, transac.Id)
	// } else {
	// 	c.String(404, "La transaccion no existe")
	// }
}

var transacciones []Transaccion

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

	err = json.Unmarshal(data, &transacciones)

	if err != nil {
		panic("error haciendo el unmarshal")
	}

	transacciones := router.Group("/transacciones")

	transacciones.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GetAll": transacciones,
		})
	})

	transacciones.GET("/:id", buscarTransaccion)

	router.Run()
}
