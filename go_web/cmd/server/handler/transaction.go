package handler

import "github.com/gin-gonic/gin"

type request struct {
	Transaction_Code string  `form:"transaction_code", json:"transaction_code"`
	Coin             string  `form:"coin", json:"coin"`
	Amount           float64 `form:"amount", json:"amount"`
	Emitor           string  `form:"emitor", json:"emitor"`
	Receptor         string  `form:"receptor", json:"receptor"`
	Transaction_Date string  `form:"transaction_date", json:"transaction_date"`
}

type Controller struct {
	service transactions.Service
}transaction_code, coin, emitor, receptor, transaction_date string, amount float64

func NewController(ser transactions.Service) *Controller {
	return &Controller{service: ser}
}

func (contr *Controller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transactions, err := contr.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error: %v \n", err)
		} else {
			ctx.JSON(200, transactions)
		}
	}
}

func (contr *Controller) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var trans request

		err := ctx.ShouldBindJSON(&trans)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar la transaction %v", err)
		} else {
			response, err := contr.service.Store(trans.Transaction_Code, trans.Coin, trans.Emitor, trans.Receptor, trans.Transaction_Date, trans.Amount)
			if err != nil {
				ctx.String(400, "No se pudo cargar la transaction: %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}
