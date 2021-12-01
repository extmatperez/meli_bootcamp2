package handler

import (
	transacciones "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transacciones"
	"github.com/gin-gonic/gin"
)


type request struct {
	CodigoTransaccion int     `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

type Transaccion struct {
	service transacciones.Service
}

func NewTransaccion(t transacciones.Service) *Transaccion {
	return &Transaccion{service: t}
}

func (t Transaccion)GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
			data, err := t.service.GetAll()
			if err != nil{
			ctx.String(200, "Ocurrio un error")
			return
			}
			ctx.JSON(200, data)
		}
	}


func (t *Transaccion) Store() gin.HandlerFunc{
	return func(ctx *gin.Context){
	var req request
	err := ctx.ShouldBind(&req)
	token := ctx.GetHeader("token")

	if token == "secure" {
		if err != nil {
			ctx.String(400, "Ha ocurrido un error")
		} else {
			p, err := t.service.Store(req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)
			if err != nil {
				ctx.String(400, "Ha ocurrido un error al escribir")
				return
			}
			ctx.JSON(200, p)
		}
	} else {
		ctx.String(401, "No tiene permisos para realizar la petición realizada")
	}
}

}




