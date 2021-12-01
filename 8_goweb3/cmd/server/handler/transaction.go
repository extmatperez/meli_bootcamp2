package handler

import (
	"fmt"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/internal/transactions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type request struct {
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Remitter string  `json:"remitter"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(ser transactions.Service) *Transaction {
	return &Transaction{service: ser}
}

func (transact *Transaction) GetAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		transactions, err := transact.service.GetAll()
		if err != nil {
			context.String(400, "Hubo un error")
		} else {
			context.JSON(200, transactions)
		}
	}
}

func (transact *Transaction) Store() gin.HandlerFunc {
	return func(context *gin.Context) {
		var newTransaction request
		err := context.ShouldBindJSON(&newTransaction)
		if err != nil {
			context.String(400, "Hubo un error")
			return
		}
		transactionCreated, err := transact.service.Store(0, newTransaction.Code, newTransaction.Currency,
			newTransaction.Amount, newTransaction.Remitter, newTransaction.Receptor, newTransaction.Date)
		if err != nil {
			context.String(400, "Hubo un error")
		} else {
			context.JSON(200, transactionCreated)
		}
	}
}

func (transact *Transaction) Update() gin.HandlerFunc {
	return func(context *gin.Context) {
		var modTransaction request
		err := context.ShouldBindJSON(&modTransaction)
		if err != nil {
			context.String(400, "Hubo un error")
			return
		}
		idStr := context.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			context.String(400, "Hubo un error")
			return
		}
		transactionMod, err := transact.service.Update(id, modTransaction.Code, modTransaction.Currency,
			modTransaction.Amount, modTransaction.Remitter, modTransaction.Receptor, modTransaction.Date)
		if err != nil {
			context.String(404, err.Error())
		} else {
			context.JSON(200, transactionMod)
		}
	}
}

func (transact *Transaction) Delete() gin.HandlerFunc {
	return func(context *gin.Context) {
		idStr := context.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			context.String(400, "Hubo un error")
			return
		}
		err = transact.service.Delete(id)
		if err != nil {
			context.String(404, err.Error())
		} else {
			context.JSON(200, fmt.Sprintf("La transaccion con id %v, se ha eliminado correctamente", id))
		}
	}
}

func (transact *Transaction) ModifyTransactionCode() gin.HandlerFunc {
	return func(context *gin.Context) {
		idStr := context.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			context.String(400, "Hubo un error")
			return
		}
		transactionCode := context.Param("TransactionCode")
		transactionMod, err := transact.service.ModifyTransactionCode(id, transactionCode)
		if err != nil {
			context.String(404, err.Error())
		} else {
			context.JSON(200, transactionMod)
		}
	}
}

func (transact *Transaction) ModifyAmount() gin.HandlerFunc {
	return func(context *gin.Context) {
		idStr := context.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			context.String(400, "Hubo un error")
			return
		}
		amountStr := context.Param("Amount")
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			context.String(400, "Hubo un error")
			return
		}
		transactionMod, err := transact.service.ModifyAmount(id, amount)
		if err != nil {
			context.String(404, err.Error())
		} else {
			context.JSON(200, transactionMod)
		}
	}
}