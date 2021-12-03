package handler

import (
	"strconv"

	transaction "github.com/extmatperez/meli_bootcamp2/7_goweb2/turn_afternoon/internal/transaction"
	"github.com/gin-gonic/gin"
)

type request struct {
	TransactionCode string  `json:"transaction_code"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Receiver        string  `json:"receiver"`
	Sender          string  `json:"sender"`
	TransactionDate string  `json:"transaction_date"`
}

type Transaction struct {
	service transaction.Service
}

func NewTransaction(s transaction.Service) *Transaction {
	return &Transaction{service: s}
}

func (tran *Transaction) GetAll() gin.HandlerFunc {

	return func(context *gin.Context) {
		token := context.GetHeader("token")

		if token != "" {
			if token == "dig.123" {
				transactions, err := tran.service.GetAll()
				if err != nil {
					context.String(400, "There was a error %v", err)
				} else {
					context.JSON(200, gin.H{
						"transaction": transactions})
				}

			} else {
				context.String(401, "Invalid Token")
			}
		} else {
			context.String(400, "Need enter a token")
		}

	}
}

func (tran *Transaction) Store() gin.HandlerFunc {

	return func(context *gin.Context) {
		token := context.GetHeader("token")

		if token != "" {
			if token == "dig.123" {
				var newTransaction request
				err := context.ShouldBindJSON(&newTransaction)
				if err != nil {
					context.String(400, "Hubo un error al querer cargar una transaction %v", err)
				} else {
					response, err := tran.service.Store(newTransaction.TransactionCode, newTransaction.Currency, newTransaction.Amount,
						newTransaction.Receiver, newTransaction.Sender, newTransaction.TransactionDate)
					if err != nil {
						context.String(400, "No se pudo cargar la persona %v", err)
					} else {
						context.JSON(200, response)
					}
				}

			} else {
				context.String(401, "Invalid Token")
			}
		} else {
			context.String(400, "Need enter a token")
		}
	}
}

func (tran *Transaction) GetByID() gin.HandlerFunc {

	return func(context *gin.Context) {
		token := context.GetHeader("token")

		if token != "" {
			if token == "dig.123" {

				idParam, err := strconv.Atoi(context.Param("id"))
				if err != nil {
					context.JSON(400, gin.H{
						"transaction": err})
				}
				gotTrans, err := tran.service.GetByID(idParam)
				if err != nil {
					context.String(404, "Transaction %v not found", idParam)
				} else {
					context.JSON(200, gin.H{
						"transaction": gotTrans})

				}

			} else {
				context.String(401, "Invalid Token")
			}
		} else {
			context.String(400, "Need enter a token")
		}
	}
}

func (trans *Transaction) GetByReceiver() gin.HandlerFunc {

	return func(context *gin.Context) {

		token := context.GetHeader("token")

		if token != "" {
			if token == "dig.123" {
				var receiver string = context.Param("receiver")
				transFound, err := trans.service.GetByReceiver(receiver)
				if err != nil {
					context.String(404, "Receiver %s not found, error: %v", receiver, err)

				} else {
					context.JSON(200, gin.H{
						"transaction": transFound})

				}
			} else {
				context.String(401, "Invalid Token")
			}
		} else {
			context.String(400, "Need enter a token")
		}

	}

}
