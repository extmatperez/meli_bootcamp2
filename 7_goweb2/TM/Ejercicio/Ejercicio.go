// La diferencia de este archivo con respecto al de Go Web 1 es que la variable que tiene todas las transacciones es una variable global que se llena con un
// metodo Load, que si no no tiene nada y tiene que ser lo que primero se corra para poder tener resultados.
// Falta validacion sobre POST de Payment!
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var all_payments []Payment

type Payment struct {
	Id       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"fecha"`
}

func add_payment(ctx *gin.Context) {
	var pay Payment
	err := ctx.ShouldBindJSON(&pay)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		if len(all_payments) == 0 {
			pay.Id = 1
		} else {
			pay.Id = all_payments[len(all_payments)-1].Id + 1
		}
		all_payments = append(all_payments, pay)
		ctx.JSON(200, pay)
	}
}

func load_payments(ctx *gin.Context) {
	bytes, err := os.ReadFile("/Users/rovega/Documents/GitHub/meli_bootcamp2/7_goweb2/payment.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "No se han podido cargar las transacciones",
		})
		return
	}

	errUnmarshal := json.Unmarshal(bytes, &all_payments)

	if errUnmarshal != nil {
		ctx.JSON(500, gin.H{
			"error": "No se pudo parsear correctamente el JSON de transacciones.",
		})
	} else {
		ctx.String(200, "Transacciones cargadas.")
	}
}

func filtrar(slicePayments []Payment, field string, value string) []Payment {
	var filtered []Payment

	var pay Payment
	types := reflect.TypeOf(pay)
	i := 0
	for i = 0; i < types.NumField(); i++ {
		if strings.ToLower(types.Field(i).Name) == field {
			break
		}
	}

	for _, v := range slicePayments {
		var cadena string
		cadena = fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, value) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func getAllPayments(ctx *gin.Context) {
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

func filtrar_payments(ctx *gin.Context) {
	var etiquetas []string
	etiquetas = append(etiquetas, "moneda", "emisor", "receptor")

	var filtered []Payment
	filtered = all_payments

	for _, v := range etiquetas {
		if len(ctx.Query(v)) != 0 && len(filtered) != 0 {
			filtered = filtrar(filtered, v, ctx.Query(v))
		}
	}

	if len(filtered) == 0 {
		ctx.String(200, "No hay coincidencias")
	} else {
		ctx.JSON(200, filtered)
	}
}

func getPayments(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != "" {
		if token == "123456" {
			if len(all_payments) > 0 {
				ctx.JSON(200, all_payments)
			} else {
				ctx.String(200, "No hay transacciones cargadas.")
			}
		} else {
			ctx.String(401, "Token incorrecto.")
		}
	} else {
		ctx.String(400, "No ingreso un token")
	}
}

func main() {
	router := gin.Default()
	payments := router.Group("/payments")

	payments.POST("/add", add_payment)
	payments.GET("/", getPayments)
	payments.GET("/load", load_payments)
	payments.GET("/filter", getAllPayments)

	router.Run()
}
