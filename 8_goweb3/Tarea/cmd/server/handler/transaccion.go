package handler

import (
	transacciones "github.com/extmatperez/meli_bootcamp2/tree/patricio_pallua/8_goweb3/Tarea/internal/Transacciones"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
}

type Transaccion struct {
	service transacciones.Service
}

func NewTransaccion(s transacciones.Service) *Transaccion {
	return &Transaccion{
		service: s,
	}
}

func (t *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234567" {
			ctx.JSON(401, gin.H{"error": "Token invalido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		t, err := t.service.Store(10, req.Codigo, req.Moneda, req.Monto, req.Emisor, req.Receptor)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, t)
	}
}
