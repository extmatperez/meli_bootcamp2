package handler

import (
	"strconv"

	productos "github.com/extmatperez/meli_bootcamp2/Semana-2/7_goweb2/proyecto/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

type Producto struct {
	service productos.Service
}

func NewProducto(ser productos.Service) *Producto {
	return &Producto{service: ser}
}

func (prod *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := prod.service.GetAll()

		if err != nil {
			ctx.JSON(400, err)
			return
		}
		ctx.JSON(200, response)
	}
}

func (prod *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(400, err)
			return
		}
		response, err := prod.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)

		if err != nil {
			ctx.JSON(400, err)
			return
		}
		ctx.JSON(200, response)
	}
}

func (prod *Producto) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(400, err)
			return
		}
		myId, err := strconv.Atoi(ctx.Param("id"))

		response, err := prod.service.Update(myId, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaCreacion)

		if err != nil {
			ctx.JSON(400, err)
			return
		}
		ctx.JSON(200, response)
	}
}

func (prod *Producto) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request
		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(400, err)
			return
		}
		myId, err := strconv.Atoi(ctx.Param("id"))

		response, err := prod.service.UpdateName(myId, req.Nombre)

		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		ctx.JSON(200, response)
	}
}

func (prod *Producto) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		myId, err := strconv.Atoi(ctx.Param("id"))

		response, err := prod.service.Delete(myId)

		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		ctx.JSON(200, response)
	}
}
