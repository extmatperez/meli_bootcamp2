package handler

import (
	"fmt"
	"os"
	"strconv"

	payments "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/internal/payments"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/pkg/web"
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

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("TOKEN")

	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "Falta token."))
		return false
	}

	token_env := os.Getenv("TOKEN")

	if token != token_env {
		ctx.JSON(404, web.NewResponse(404, nil, "Token incorrecto."))
		return false
	}

	return true
}

func (controller *Payment) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		payments, err := controller.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al cargar todas las transacciones: %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, payments, ""))
		}
	}
}

func (controller *Payment) Filter() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var pay request

		err := ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Error en el body."))
		} else {
			paymentUpdated, err := controller.service.Filter(pay.Codigo, pay.Moneda, pay.Emisor, pay.Receptor, pay.Fecha, pay.Monto)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewResponse(200, paymentUpdated, ""))
			}
		}
	}
}

func (controller *Payment) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

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
				ctx.JSON(200, web.NewResponse(200, response, ""))
			}
		}
	}
}

func (controller *Payment) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

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

func (controller *Payment) UpdateCodigo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

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

func (controller *Payment) UpdateMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

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

func (controller *Payment) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

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
