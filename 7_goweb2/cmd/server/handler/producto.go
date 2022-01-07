package handler

import (
	"strconv"

	internalProducto "github.com/extmatperez/meli_bootcamp2/tree/castiglione_adrian/7_goweb2/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre    string  `json:"nombre" binding:"required"`
	Color     string  `json:"color" binding:"required"`
	Precio    float64 `json:"precio" binding:"required"`
	Stock     int     `json:"stock" binding:"required"`
	Codigo    string  `json:"codigo" binding:"required"`
	Publicado bool    `json:"publicado" binding:"required"`
}

type Response struct {
	Code  int
	Data  interface{}
	Error string
}

type Producto struct {
	service internalProducto.Service
}

func NewProducto(s internalProducto.Service) *Producto {
	return &Producto{
		service: s,
	}
}

func (p *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productos, err := p.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, productos)
		}
	}
}

func (p *Producto) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		}

		producto, err := p.service.GetOne(id)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, producto)
		}
	}
}

func (p *Producto) AddOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		err := ctx.ShouldBind(&req)

		if err != nil {
			ctx.JSON(400, map[string]string{"msg": err.Error()})
			return
		}

		product, err := p.service.AddOne(
			req.Nombre,
			req.Codigo,
			req.Precio,
			req.Stock,
			req.Codigo,
			req.Publicado,
		)

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		}

		ctx.JSON(200, product)
	}
}
