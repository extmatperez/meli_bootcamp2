package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	ID       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
}

var transaccionesTodas []Transaccion

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
		var transferencias []Transaccion
		json.Unmarshal(res, &transferencias)

		var finalText string

		for _, t := range transferencias {
			newText := fmt.Sprintf("\n La transaccion %v por un monto de %v %v la genero %v para %v \n", t.ID, t.Monto, t.Moneda, t.Emisor, t.Receptor)
			finalText = finalText + newText
		}
		c.String(200, finalText)

	}
}

func getOne(c *gin.Context) {
	ID := c.Param("id")
	fmt.Println("EL ID ES", ID)

	res, err := os.ReadFile("./Transacciones.json")
	if err != nil {
		c.JSON(400, gin.H{
			"Mensaje": "Hubo un problema con el archivo",
		})
	} else {
		var transferencias []Transaccion
		json.Unmarshal(res, &transferencias)

		var transferencia Transaccion

		for _, t := range transferencias {
			if ID == strconv.FormatInt(int64(t.ID), 10) {
				transferencia = t
				break
			}
		}
		if transferencia.Codigo == "" {
			c.JSON(404, gin.H{
				"Mensaje": "No se encontro la transferencia",
			})
		} else {
			c.JSON(http.StatusFound, gin.H{
				"Mensaje":     "Transaccion encontrada con exito",
				"Transaccion": transferencia,
			})
		}

	}
}

func agregarTransaccion(c *gin.Context) {

	token := c.GetHeader("token")

	if token != "token_secreto" {
		c.JSON(400, gin.H{
			"error": "Aca no podes pasar!!!",
		})
	} else {
		var newTran Transaccion
		fmt.Println("Empecemos a trabajar")
		err := c.ShouldBind(&newTran)
		fmt.Println(newTran)
		if err != nil {
			fmt.Println("Entramos aca")
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Println(newTran)
			var falla string
			switch {
			case newTran.Codigo == "":
				falla = "Codigo"
			case newTran.Moneda == "":
				falla = "Moneda"
			case newTran.Monto == 0:
				falla = "Monto"
			case newTran.Emisor == "":
				falla = "Emisor"
			case newTran.Receptor == "":
				falla = "Receptor"
			default:
				falla = "Todo bien"
			}

			if falla == "Todo bien" {
				if len(transaccionesTodas) == 0 {
					newTran.ID = 1
				} else {
					newTran.ID = len(transaccionesTodas) + 1
				}
				transaccionesTodas = append(transaccionesTodas, newTran)
				fmt.Println(transaccionesTodas)
				c.JSON(200, transaccionesTodas)
			} else {
				c.JSON(400, gin.H{
					"error": "Falta " + falla,
				})
			}
		}
	}

}

func main() {

	router := gin.Default()

	router.GET("/saludo", saludar)
	router.GET("/saludoLindo/:nombre", saludar2)

	transacciones := router.Group("/transacciones")
	{
		transacciones.GET("/lindo", getAllLindo)
		transacciones.GET("/feo", getAllFeo)
		transacciones.GET("/:id", getOne)
		transacciones.POST("/", agregarTransaccion)
	}

	router.Run()

}
