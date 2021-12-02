package handler

import (
	"strconv"

	payments "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/internal/payments"
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

func (controller *Payment) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "" {
			if token == "123456" {
				payments, err := controller.service.GetAll()

				if err != nil {
					ctx.String(400, "Hubo un error al cargar todas las transacciones: %v", err)
				} else {
					ctx.JSON(200, payments)
				}
			}
		} else {
			ctx.String(400, "No ingreso un token")
		}
	}
}

func (controller *Payment) Filter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "" {
			if token == "123456" {
				var pay request

				err := ctx.ShouldBindJSON(&pay)

				if err != nil {
					ctx.String(400, "Error en el body.")
				} else {
					paymentUpdated, err := controller.service.Filter(pay.Codigo, pay.Moneda, pay.Emisor, pay.Receptor, pay.Fecha, pay.Monto)
					if err != nil {
						ctx.JSON(400, err.Error())
					} else {
						ctx.JSON(200, paymentUpdated)
					}
				}
			}
		}
	}
}

func (controller *Payment) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pay request

		err := ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una transacción: %v", err)
		} else {
			response, err := controller.service.Store(pay.Codigo, pay.Moneda, pay.Emisor, pay.Receptor, pay.Fecha, pay.Monto)

			if err != nil {
				ctx.String(400, "No se pudo cargar la transacción: %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

func (controller *Payment) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pay request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(404, "El id es invalido y no fue encontrado.")
		}

		err = ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.String(400, "Error en el body.")
		} else {
			paymentUpdated, err := controller.service.Update(id, pay.Codigo, pay.Moneda, pay.Emisor, pay.Receptor, pay.Fecha, pay.Monto)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, paymentUpdated)
			}
		}
	}

}

func (controller *Payment) UpdateCodigo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pay request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(404, "El id es invalido y no fue encontrado.")
		}

		err = ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.String(400, "Error en el body.")
		} else {
			if pay.Codigo == "" {
				ctx.String(404, "El código ingresado no puede ser vacio.")
				return
			}
			paymentUpdated, err := controller.service.UpdateCodigo(id, pay.Codigo)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, paymentUpdated)
			}
		}
	}
}

func (controller *Payment) UpdateMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pay request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(404, "El id es invalido y no fue encontrado.")
		}

		err = ctx.ShouldBindJSON(&pay)

		if err != nil {
			ctx.String(400, "Error en el body.")
		} else {
			if pay.Monto == 0.0 {
				ctx.String(404, "El monto ingresado no puede ser cero.")
				return
			}
			paymentUpdated, err := controller.service.UpdateMonto(int(id), pay.Monto)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, paymentUpdated)
			}
		}
	}
}

func (controller *Payment) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(404, "El id es invalido y no fue encontrado.")
		}

		msg, err := controller.service.Delete(id)
		if err != nil {
			ctx.JSON(400, err.Error())
		} else {
			ctx.String(200, msg)
		}
	}
}
