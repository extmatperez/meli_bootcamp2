package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID                  int     `json:"id"`
	CodigoDeTransaccion string  `json:"codigo_de_transaccion"`
	Moneda              string  `json:"moneda"`
	Monto               float64 `json:"monto"`
	Emisor              string  `json:"emisor"`
	Receptor            string  `json:"receptor"`
	FechaDeTransaccion  string  `json:"fecha_de_transaccion"`
}

func sayHello(c *gin.Context) {
	var status int
	message := ""
	messageName := "message"

	if name, ok := c.GetQuery("nombre"); ok {
		if name != "" {
			status = 200
			message = "¡Hola, " + name + "!"

		} else {
			status = 400
			message = "'nombre' cannot be empty"
			messageName = "error"
		}
	} else {
		status = 400
		message = "param 'nombre' not received"
		messageName = "error"
	}

	c.JSON(status, gin.H{
		messageName: message,
	})

}

func GetAll(c *gin.Context) {
	var status int
	var response interface{}

	data, err := os.ReadFile("./transactions.json")

	if err != nil {
		status = http.StatusInternalServerError

	} else {
		err = json.Unmarshal(data, &response)
		if err != nil {
			status = http.StatusInternalServerError
		} else {
			status = http.StatusOK
			c.JSON(status, response)
		}
	}
	if status != http.StatusOK {
		c.JSON(status, gin.H{
			"error": "could not found products",
		})
	}

}

func main() {

	router := gin.Default()

	router.GET("/hello", sayHello)
	router.GET("/transactions", GetAll)

	router.Run()

}
