package handler

import (
	"fmt"
	"os"
	"strconv"

	transaction "github.com/extmatperez/meli_bootcamp2/12_testing3/turn_afternoon/internal/transaction"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/pkg/web"
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

func validToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "Invalid token"))
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.JSON(404, web.NewResponse(404, nil, "Incorrect token"))
		return false
	}

	return true
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (tran *Transaction) GetAll() gin.HandlerFunc {

	return func(context *gin.Context) {
		if !validToken(context) {
			return
		}
		transactions, err := tran.service.GetAll()
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was a error %v", err)))
		} else {
			context.JSON(200, web.NewResponse(200, transactions, ""))
		}

	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (tran *Transaction) Store() gin.HandlerFunc {

	return func(context *gin.Context) {
		if !validToken(context) {
			return
		}
		var newTransaction request
		err := context.ShouldBindJSON(&newTransaction)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was a error %v", err)))
		} else {
			response, err := tran.service.Store(newTransaction.TransactionCode, newTransaction.Currency, newTransaction.Amount,
				newTransaction.Receiver, newTransaction.Sender, newTransaction.TransactionDate)
			if err != nil {
				context.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Imposible create a transaction: %v", err)))
			} else {
				context.JSON(200, web.NewResponse(200, response, ""))
			}
		}

	}
}

/* func (tran *Transaction) CreateTransaction() gin.HandlerFunc {

	return func(context *gin.Context) {
		token := context.GetHeader("token")

		if token != "" {
			if token == os.Getenv("TOKEN") {
				var request request
				err := context.ShouldBindJSON(&request)
				if err != nil {
					context.String(400, "Hubo un error al querer cargar una transaction %v", err)
				} else {
					newTransaction := transaction.Transaction{0, request.TransactionCode, request.Currency, request.Amount,
						request.Receiver, request.Sender, request.TransactionDate}
					request, err := tran.service.CreateTransaction(newTransaction)
					if err != nil {
						context.String(400, "No se pudo cargar la transaction %v", err)
					} else {
						context.JSON(200, request)
					}
				}

			} else {
				context.String(401, "Invalid Token")
			}
		} else {
			context.String(400, "Need enter a token")
		}
	}
} */

func (tran *Transaction) GetByID() gin.HandlerFunc {

	return func(context *gin.Context) {
		if !validToken(context) {
			return
		}

		idParam, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was a error change string to id int: %v", err)))
		}
		gotTrans, err := tran.service.GetByID(idParam)
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("Transaction %v not found", idParam)))
		} else if gotTrans.ID == 0 {
			context.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("Transaction %v not found", idParam)))
		} else {
			context.JSON(200, web.NewResponse(200, gotTrans, ""))

		}

	}
}

func (trans *Transaction) GetByReceiver() gin.HandlerFunc {

	return func(context *gin.Context) {

		if !validToken(context) {
			return
		}
		var receiver string = context.Param("receiver")
		transFound, err := trans.service.GetByReceiver(receiver)
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("Receiver %s not found, error: %v", receiver, err)))

		} else {
			context.JSON(200, web.NewResponse(200, transFound, ""))

		}

	}

}

func (trans *Transaction) UpdateTransaction() gin.HandlerFunc {

	return func(context *gin.Context) {

		if !validToken(context) {
			return
		}
		id, err1 := strconv.Atoi(context.Param("id"))
		if err1 != nil {
			context.JSON(400, gin.H{
				"transaction": err1})
		}
		var request request
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was a error %v", err)))
		}
		transFound, err := trans.service.UpdateTransaction(id, request.TransactionCode, request.Currency, request.Amount,
			request.Receiver, request.Sender, request.TransactionDate)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Could update the transaction: %v", err)))

		} else {
			context.JSON(200, web.NewResponse(200, transFound, ""))

		}

	}

}

func (trans *Transaction) UpdateAmount() gin.HandlerFunc {

	return func(context *gin.Context) {

		if !validToken(context) {
			return
		}
		id, err1 := strconv.Atoi(context.Param("id"))
		if err1 != nil {
			context.JSON(400, gin.H{
				"transaction": err1})
		}
		amount, err := strconv.ParseFloat(context.Param("amount"), 32)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, "Error to parse string to float64"))

		}
		transFound, err := trans.service.UpdateAmount(id, amount)
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("Receiver %s not found, error: %v", id, err)))

		} else {
			context.JSON(200, web.NewResponse(200, transFound, ""))

		}
	}
}

func (trans *Transaction) DeleteTransaction() gin.HandlerFunc {

	return func(context *gin.Context) {

		if !validToken(context) {
			return
		}
		id, err1 := strconv.Atoi(context.Param("id"))
		if err1 != nil {
			context.JSON(400, web.NewResponse(400, nil, "There was a error"))
		}
		err := trans.service.DeleteTransaction(id)
		if err != nil {
			context.JSON(404, web.NewResponse(400, nil, "There was a error"))

		} else {
			context.JSON(200, web.NewResponse(200, "Deleted", ""))
		}
	}
}
