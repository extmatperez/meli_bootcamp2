package handler

import (
	"errors"
	"net/http"
	"strconv"

	transactions "github.com/extmatperez/meli_bootcamp2/tree/soto_jose/8_goweb3/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	Code     string `json:"code"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Date     string `json:"date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(ser transactions.Service) *Transaction {
	return &Transaction{service: ser}
}

func (per *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := per.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, transactions)
		}
	}
}

func (controller *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var t request

		err := ctx.ShouldBindJSON(&t)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Store(t.Code, t.Currency, t.Amount, t.Sender, t.Receiver, t.Date)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (controller *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body request

		err := ctx.ShouldBindJSON(&body)
		paramId := ctx.Param("id")
		id, parseErr := strconv.Atoi(paramId)

		if err != nil || parseErr != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		response, err := controller.service.Update(id, body.Code, body.Currency, body.Amount, body.Sender, body.Receiver, body.Date)
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
		} else {
			ctx.JSON(200, response)
		}

	}
}

func (controller *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		paramId := ctx.Param("id")
		id, err := strconv.Atoi(paramId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, errors.New("not valid id"))
			return
		}

		err = controller.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, errors.New("transaction not found"))
			return
		}
		ctx.Writer.WriteHeader(http.StatusNoContent)
	}
}
