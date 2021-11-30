package main

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID               int     `form:"id", json:"id"`
	Transaction_Code string  `form:"transaction_code", json:"transaction_code", binding:"required"`
	Coin             string  `form:"coin", json:"coin", binding:"required"`
	Amount           float64 `form:"amount", json:"amount", binding:"required"`
	Emitor           string  `form:"emitor", json:"emitor", binding:"required"`
	Receptor         string  `form:"receptor", json:"receptor", binding:"required"`
	Transaction_Date string  `form:"transaction_date", json:"transaction_date", binding:"required"`
}

var ListaTransactions []Transaction

func addTransaction(ctx *gin.Context) {
	var transaction Transaction

	token := ctx.GetHeader("token")

	if token == "123456" {
		err := ctx.ShouldBindJSON(&transaction)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			validFields := fieldsValidator(transaction)
			if validFields == "ok" {
				var id = len(ListaTransactions) + 1
				transaction.ID = id
				ListaTransactions = append(ListaTransactions, transaction)
				ctx.JSON(200, transaction)
			} else {
				ctx.String(400, "el campo %s es requerido", validFields)
			}
		}
	} else {
		ctx.String(401, "no tiene permisos para realizar la petición solicitada")
	}
}

func fieldsValidator(transaction Transaction) string {
	typesOfFields := reflect.ValueOf(transaction)

	i := 0
	for i = 0; i < typesOfFields.NumField(); i++ {
		valueOfField := typesOfFields.Field(i).Interface()
		if valueOfField == "" {
			field := reflect.TypeOf(transaction).Field(i).Name
			return field
		}
	}
	return "ok"
}

func main() {
	r := gin.Default()

	r.POST("/agregarEntidad", addTransaction)

	r.Run()
}
