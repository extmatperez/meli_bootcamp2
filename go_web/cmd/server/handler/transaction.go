package handler

import (
	"fmt"
	"os"
	"strconv"

	transactions "github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/internal/transaction"
	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/go_web/pkg/web"
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

// ListTransactions godoc
// @Summary List transactions
// @Tags Transaction
// @Description get all transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions/get [get]
func (contr *Controller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tokenValidator(ctx) {
			return
		}
		transactions, err := contr.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, transactions, ""))
		}
	}
}

// NewTransaction godoc
// @Summary Add Transaction
// @Tags Transaction
// @Description post transaction
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions/add [post]
func (contr *Controller) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tokenValidator(ctx) {
			return
		}

		var trans request

		err := ctx.ShouldBindJSON(&trans)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar la transaction %v", err)))
		} else {
			response, err := contr.service.Store(trans.Transaction_Code, trans.Coin, trans.Emitor, trans.Receptor, trans.Transaction_Date, trans.Amount)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar la transaction: %v", err)))
			} else {
				ctx.JSON(200, web.NewResponse(200, response, ""))
			}
		}
	}
}

// UpdateTransaction godoc
// @Summary Update Transaction
// @Tags Transaction
// @Description put transaction
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /transactions/:id [put]
func (controller *Controller) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !tokenValidator(ctx) {
			return
		}

		var trans request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("El id es invalido", err)))
		}

		err = ctx.ShouldBindJSON(&trans)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error en el body %v", err)))
		} else {
			transactionUpdated, err := controller.service.Update(id, trans.Transaction_Code, trans.Coin, trans.Emitor, trans.Receptor, trans.Transaction_Date, trans.Amount)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewResponse(200, transactionUpdated, ""))
			}
		}

	}
}

// PatchTransaction godoc
// @Summary Transaction
// @Tags Transaction
// @Description patch transaction receptor
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /transaction/:id [patch]
func (controller *Controller) UpdateReceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tokenValidator(ctx) {
			return
		}

		var trans request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("El id es invalido %v", err)))
		}

		err = ctx.ShouldBindJSON(&trans)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error en el body %v", err)))
		} else {
			if trans.Receptor == "" {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("El nombre no puede estar vacío %v", err)))
				return
			}
			transactionUpdated, err := controller.service.UpdateReceptor(id, trans.Receptor)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			} else {
				ctx.JSON(200, web.NewResponse(200, transactionUpdated, ""))
			}
		}

	}
}

// Transaction doc
// @Summary  transaction
// @Tags Transaction
// @Description delete transaction
// @Accept  json
// @Produce  json
// @Param tid path int true "id"
// @Success 200 {object} web.Response
// @Router /transaction/:id [delete]
func (controller *Controller) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !tokenValidator(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))

			ctx.String(400, "El id es invalido")
		}

		err = controller.service.Delete(id)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("La trasferencia %d ha sido eliminada", id), ""))
		}

	}
}
