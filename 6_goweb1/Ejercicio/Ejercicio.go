// 1) Elegi Transacciones
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Payment struct {
	Id          int     `json:"id"`
	Code        string  `json:"codigo"`
	Coin        string  `json:"moneda"`
	Amount      float64 `json:"monto"`
	Origin      string  `json:"emisor"`
	Destination string  `json:"receptor"`
	Date        string  `json:"fecha"`
}

func getAllPayments(ctx *gin.Context) {
	bytes, err := os.ReadFile("/Users/rovega/Documents/GitHub/meli_bootcamp2/6_goweb1/payment.json")

	if err != nil {
		defer func() {
			ctx.JSON(500, gin.H{
				"error": "No se han podido cargar las transacciones",
			})
			fmt.Println("Ejecución finalizada!")
		}()
		return
	}

	var payments []Payment
	json.Unmarshal(bytes, &payments)

	ctx.JSON(200, gin.H{
		"payments": payments,
	})

}

func main() {
	router := gin.Default()

	router.GET("/hello-world/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, fmt.Sprintf("Hola %s!", name))
	})

	router.GET("/payments", getAllPayments)

	router.Run()
}
