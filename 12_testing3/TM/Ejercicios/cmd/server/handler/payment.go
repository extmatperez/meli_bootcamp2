package handler

import (
	"fmt"
	"strconv"

	payments "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/12_testing3/TM/Ejercicios/internal/payments"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/12_testing3/TM/Ejercicios/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Codigo   string  `json:"codigo"`
	Monto    float64 `json:"monto"`
	Moneda   string  `json:"moneda"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"fecha"`
}

type Payment struct {
	service payments.Service
}

func NewPayment(s payments.Service) *Payment {
	return &Payment{service: s}
}

// ListPayments godoc
// @Summary List payments
// @Tags Payments
// @Description get all the payments
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /payments [get]
func (controller *Payment) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payments, err := controller.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al cargar todas las transacciones: %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, payments, ""))
		}
	}
}

// StorePayments godoc
// @Summary Store payments
// @Tags Payments
// @Description store payments
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param payment body request true "Payment to store"
// @Success 200 {object} web.response
// @Router /payments [post]
func (controller *Payment) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pay request

		err := ctx.ShouldBindJSON(&pay)

		// Aca es donde se deberia validar la correspondencia de los campos.
		if pay.Codigo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Es necesario ingresar el código de la transacción."))
			return
		}
		if pay.Moneda == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Es necesario ingresar la moneda."))
			return
		}
		if pay.Emisor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Es necesario ingresar el emisor."))
			return
		}
		if pay.Receptor == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Es necesario ingresar el receptor."))
			return
		}
		if pay.Fecha == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Es necesario ingresar la fecha."))
			return
		}
		if pay.Monto == 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Es necesario que el monto sea diferente de cero."))
			return
		}

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una transacción: %v", err)
		} else {
			response, err := controller.service.Store(pay.Codigo, pay.Moneda, pay.Emisor, pay.Receptor, pay.Fecha, pay.Monto)

			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar la transacción: %v", err)))
			} else {
				ctx.JSON(201, web.NewResponse(201, response, ""))
			}
		}
	}
}

// UpdatePayments godoc
// @Summary Update payments
// @Tags Payments
// @Description update some fields of one payment
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /payments/:id [put]
func (controller *Payment) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pay request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "El id es inválido."))
		}

		err = ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Error en el body."))
		} else {
			paymentUpdated, err := controller.service.Update(id, pay.Codigo, pay.Moneda, pay.Emisor, pay.Receptor, pay.Fecha, pay.Monto)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewResponse(200, paymentUpdated, ""))
			}
		}
	}

}

// PatchPaymentCode godoc
// @Summary Update payment code
// @Tags Payments
// @Description change one payment's code
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /payments/code/:id [patch]
func (controller *Payment) UpdateCodigo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pay request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "El id es inválido."))
		}

		err = ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Error en el body."))
		} else {
			if pay.Codigo == "" {
				ctx.JSON(404, web.NewResponse(404, nil, "El código ingresado no puede ser vacio."))
				return
			}
			paymentUpdated, err := controller.service.UpdateCodigo(id, pay.Codigo)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewResponse(200, paymentUpdated, ""))
			}
		}
	}
}

// PatchPaymentAmount godoc
// @Summary Update payment amount
// @Tags Payments
// @Description change one payment's amount
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /payments/amount/:id [patch]
func (controller *Payment) UpdateMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pay request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "El id es inválido."))
		}

		err = ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Error en el body."))
		} else {
			if pay.Monto == 0.0 {
				ctx.JSON(404, web.NewResponse(404, nil, "El monto ingresado no puede ser cero."))
				return
			}
			paymentUpdated, err := controller.service.UpdateMonto(int(id), pay.Monto)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewResponse(200, paymentUpdated, ""))
			}
		}
	}
}

// DeletePayment godoc
// @Summary Delete payment
// @Tags Payments
// @Description delete one payment
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.response
// @Router /payments/:id [delete]
func (controller *Payment) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "El id es inválido."))
		}

		msg, err := controller.service.Delete(id)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewResponse(200, msg, ""))
		}
	}
}
