// 1) Elegi Transacciones
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Payment struct {
	Id       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"fecha"`
}

func getAllPayments(ctx *gin.Context) {
	bytes, err := os.ReadFile("/Users/rovega/Documents/GitHub/meli_bootcamp2/6_goweb1/payment.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "No se han podido cargar las transacciones",
		})
		return
	}

	var all_payments []Payment
	errUnmarshal := json.Unmarshal(bytes, &all_payments)

	if errUnmarshal != nil {
		ctx.JSON(500, gin.H{
			"error": "No se pudo parsear correctamente el JSON de transacciones.",
		})
	}

	filter_code := ctx.Query("codigo")
	filter_coin := ctx.Query("moneda")
	filter_amount := ctx.Query("monto")
	filter_origin := ctx.Query("emisor")
	filter_destination := ctx.Query("receptor")
	filter_date := ctx.Query("fecha")

	var filtered_payments []Payment

	for i := 0; i < len(all_payments); i++ {
		is_filtered := true

		if filter_code != "" {
			if !strings.EqualFold(all_payments[i].Codigo, filter_code) {
				is_filtered = false
			}
		}

		if filter_coin != "" {
			if !strings.EqualFold(all_payments[i].Moneda, filter_coin) {
				is_filtered = false
			}
		}

		if filter_amount != "" {
			check, err := strconv.ParseFloat(filter_amount, 64)
			if err == nil {
				if all_payments[i].Monto != check {
					is_filtered = false
				}
			}
		}

		if filter_origin != "" {
			if !strings.EqualFold(all_payments[i].Emisor, filter_origin) {
				is_filtered = false
			}
		}

		if filter_destination != "" {
			if !strings.EqualFold(all_payments[i].Receptor, filter_destination) {
				is_filtered = false
			}
		}

		if filter_date != "" {
			if !strings.EqualFold(all_payments[i].Fecha, filter_date) {
				is_filtered = false
			}
		}

		// Chequea los filtros si los cumple todos.
		if is_filtered {
			filtered_payments = append(filtered_payments, all_payments[i])
		}
	}

	ctx.JSON(200, gin.H{
		"payments": filtered_payments,
	})

}

func getPayment(ctx *gin.Context) {
	paymentId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "ID invalido",
		})
		return
	}

	bytes, err := os.ReadFile("/Users/rovega/Documents/GitHub/meli_bootcamp2/6_goweb1/payment.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "No se han podido cargar las transacciones",
		})
		return
	}

	var all_payments []Payment
	errUnmarshal := json.Unmarshal(bytes, &all_payments)

	if errUnmarshal != nil {
		ctx.JSON(500, gin.H{
			"error": "Error al parsear JSON de transacciones",
		})
		return
	}

	var payment Payment

	for i := 0; i < len(all_payments); i++ {
		if all_payments[i].Id == paymentId {
			payment = all_payments[i]
			break
		}
	}

	if payment == (Payment{}) {
		ctx.JSON(404, gin.H{
			"message": "Transaccion no encontrada",
		})
		return
	}

	ctx.JSON(200, payment)
}

func main() {
	router := gin.Default()

	router.GET("/hello-world/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, fmt.Sprintf("Hola %s!", name))
	})

	payments := router.Group("/payments")
	{
		payments.GET("/", getAllPayments)
		payments.GET("/:id", getPayment)
	}

	router.Run()
}
