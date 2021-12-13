package handler

import (
	"os"
	"strconv"

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


// ListTransacciones godoc
// @Summary List transacciones
// @Tags Transacciones
// @Description get transacciones
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transacciones [get]

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


// ListTransacciones godoc
// @Summary List transacciones
// @Tags Transacciones
// @Description post transacciones
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transacciones [post]

func (t *Transaccion) Store() gin.HandlerFunc{
	return func(ctx *gin.Context){
	var req request
	err := ctx.ShouldBind(&req)
	token := ctx.GetHeader("token")

	if token == os.Getenv("TOKEN") {
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

func (t *Transaccion) Update() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "secure"{
			ctx.JSON(401, gin.H{"error": "Token invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "ID invalido"})
			return
		}
		var req request
		err = ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		t, err := t.service.Update(int(id), req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaTransaccion)
		if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
		}
		ctx.JSON(200, t)
	}
}

func (t *Transaccion) UpdateEmisor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "secure"{
			ctx.JSON(401, gin.H{"error": "Token invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "ID invalido"})
			return
		}
		var req request
		err = ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		t, err := t.service.UpdateEmisor(int(id), req.Emisor)
		if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
		}
		ctx.JSON(200, t)
	}
}


func (t *Transaccion) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "secure"{
			ctx.JSON(401, gin.H{"error": "Token invalido"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "ID invalido"})
			return
		}
		err = t.service.Delete(int(id))

		if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
		}

		ctx.JSON(200, gin.H{"data": "El producto ha sido eliminado"})
	}
}



