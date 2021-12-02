package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.String(http.StatusUnauthorized, "Falta token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.String(http.StatusUnauthorized, "Token incorrecto")
		return false
	}

	return true
}

func NewTransaction(ser transactions.Service) *Transaction {
	return &Transaction{service: ser}
}

func (per *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		transactions, err := per.service.GetAll()

		if err != nil {
			ctx.String(http.StatusBadRequest, "Hubo un error %v", err)
		} else {
			ctx.JSON(http.StatusOK, transactions)
		}
	}
}

func (controller *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var t request

		err := ctx.ShouldBindJSON(&t)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Store(t.Code, t.Currency, t.Amount, t.Sender, t.Receiver, t.Date)
			if err != nil {
				ctx.String(http.StatusBadRequest, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(http.StatusOK, response)
			}
		}

	}
}

func validateUpdatePayload(payload request) error {

	if payload.Amount == 0 || payload.Code == "" || payload.Currency == "" || payload.Date == "" || payload.Receiver == "" || payload.Sender == "" {
		return errors.New("unvalid body")
	}
	return nil
}

func (controller *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}
		var body request

		err := ctx.ShouldBindJSON(&body)
		paramId := ctx.Param("id")
		id, parseErr := strconv.Atoi(paramId)

		bodyErr := validateUpdatePayload(body)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if parseErr != nil {
			ctx.JSON(http.StatusBadRequest, parseErr.Error())
			return
		}
		if bodyErr != nil {
			ctx.JSON(http.StatusBadRequest, bodyErr.Error())
			fmt.Println(bodyErr)
			return
		}

		response, err := controller.service.Update(id, body.Code, body.Currency, body.Amount, body.Sender, body.Receiver, body.Date)
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
		} else {
			ctx.JSON(http.StatusOK, response)
		}

	}
}

func (controller *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

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

func (controller *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var body request

		err := ctx.ShouldBindJSON(&body)
		paramId := ctx.Param("id")
		id, parseErr := strconv.Atoi(paramId)

		if err != nil || parseErr != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		response, err := controller.service.UpdateCodeAndAmount(id, body.Code, body.Amount)
		if err != nil {
			ctx.AbortWithError(http.StatusNotFound, err)
		}
		ctx.JSON(http.StatusOK, response)
	}
}
