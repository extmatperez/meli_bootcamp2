package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp2/tree/soto_jose/11_testing2/TT/pkg/web"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/soto_jose/11_testing2/TT/transactions"
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

			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, fmt.Sprintf("Hubo un error %v", err)))
			return
		}
		web.NewResponse(http.StatusOK, transactions, "")
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transactions, ""))
	}
}

func (controller *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var t request

		err := ctx.ShouldBindJSON(&t)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, fmt.Sprintf("Body error: %v", err)))
			return
		}

		response, err := controller.service.Store(t.Code, t.Currency, t.Amount, t.Sender, t.Receiver, t.Date)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, fmt.Sprintf("Error saving the transaction %v", err)))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, response, ""))
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

		var body request

		err := ctx.ShouldBindJSON(&body)
		paramId := ctx.Param("id")
		id, parseErr := strconv.Atoi(paramId)

		bodyErr := validateUpdatePayload(body)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if parseErr != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, parseErr.Error()))
			return
		}
		if bodyErr != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, bodyErr.Error()))
			return
		}

		response, err := controller.service.Update(id, body.Code, body.Currency, body.Amount, body.Sender, body.Receiver, body.Date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		} else {
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, response, ""))
		}
	}
}

func (controller *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		paramId := ctx.Param("id")
		id, err := strconv.Atoi(paramId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "not valid id"))
			return
		}

		err = controller.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("Transaction with id %v not found", id)))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}

func (controller *Transaction) UpdateCodeAndAmount() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body request

		err := ctx.ShouldBindJSON(&body)
		paramId := ctx.Param("id")
		id, parseErr := strconv.Atoi(paramId)

		if err != nil || parseErr != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		}

		response, err := controller.service.UpdateCodeAndAmount(id, body.Code, body.Amount)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, response, ""))
	}
}
