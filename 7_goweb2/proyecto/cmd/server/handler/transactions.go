package handler

import (
	"net/http"
	"strconv"

	internal "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	CodigoDeTransaccion string  `json:"codigo_de_transaccion" binding:"required"`
	Moneda              string  `json:"moneda" binding:"required"`
	Monto               float64 `json:"monto" binding:"required"`
	Emisor              string  `json:"emisor" binding:"required"`
	Receptor            string  `json:"receptor" binding:"required"`
	FechaDeTransaccion  string  `json:"fecha_de_transaccion" binding:"required"`
}

type Transaction struct {
	service internal.Service
}

func NewTransaction(serv internal.Service) *Transaction {
	return &Transaction{service: serv}
}

func (t *Transaction) GetAll() gin.HandlerFunc { //TODO: implement filters
	return func(ctx *gin.Context) {
		if validateToken(ctx) {
			response, err := t.service.GetAll()
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, response)
		}
	}
}

func (t *Transaction) GetTransactionByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if validateToken(ctx) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid id",
				})
				return
			}
			response, err := t.service.GetTransactionByID(id)

			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, response)
			return
		}
	}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if validateToken(ctx) {
			var tr request
			err := ctx.ShouldBindJSON(&tr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid arguments",
				})
				return
			}
			respuesta, err := t.service.Store(tr.CodigoDeTransaccion, tr.Moneda, tr.Monto, tr.Emisor, tr.Receptor, tr.FechaDeTransaccion)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, respuesta)
			return
		}

	}
}

//Funciones auxiliares

func validateToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token != "" {
		if token == "123456" {
			return true
		}
		ctx.String(http.StatusUnauthorized, "Token incorrecto")
		return false
	}
	ctx.String(http.StatusUnauthorized, "No se ingresó un token")
	return false
}
