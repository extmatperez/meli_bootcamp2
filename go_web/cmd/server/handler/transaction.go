package handler

import (
	"os"
	"strconv"

	transactions "github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/internal/transaction"
	"github.com/gin-gonic/gin"
)

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
}

func NewController(ser transactions.Service) *Controller {
	return &Controller{service: ser}
}

func tokenValidator(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.String(400, "Missing token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.String(404, "Wrong token")
		return false
	}
	return true
}

func (contr *Controller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tokenValidator(ctx) {
			return
		}
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
		if !tokenValidator(ctx) {
			return
		}

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

func (controller *Controller) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !tokenValidator(ctx) {
			return
		}

		var trans request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&trans)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			transactionUpdated, err := controller.service.Update(id, trans.Transaction_Code, trans.Coin, trans.Emitor, trans.Receptor, trans.Transaction_Date, trans.Amount)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, transactionUpdated)
			}
		}

	}
}

func (controller *Controller) UpdateReceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tokenValidator(ctx) {
			return
		}

		var trans request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&trans)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			if trans.Receptor == "" {
				ctx.String(404, "El nombre no puede estar vacío")
				return
			}
			transactionUpdated, err := controller.service.UpdateReceptor(id, trans.Receptor)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, transactionUpdated)
			}
		}

	}
}

func (controller *Controller) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tokenValidator(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = controller.service.Delete(id)
		if err != nil {
			ctx.JSON(400, err.Error())
		} else {
			ctx.String(200, "La trasferencia %d ha sido eliminada", id)
		}

	}
}
