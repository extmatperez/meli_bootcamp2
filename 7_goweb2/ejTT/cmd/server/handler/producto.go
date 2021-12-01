package handler

import (
	"strconv"

	internal "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/7_goweb2/ejTT/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre         string  `json:"nombre"`
	Color          string  `json:"color" `
	Precio         float64 `json:"precio" `
	Stock          int     `json:"stock" `
	Codigo         string  `json:"codigo"`
	Publicado      bool    `json:"publicado"`
	Fecha_creacion string  `json:"fecha_creacion"`
}
type Producto struct {
	service internal.Service
}

func NewProducto(serv internal.Service) *Producto {
	return &Producto{service: serv}
}

func (p *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, products)
	}

}

func (p *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod request

		err := ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err,
			})
		}

		product, e := p.service.Store(prod.Nombre, prod.Color, prod.Precio, prod.Stock, prod.Codigo, prod.Publicado, prod.Fecha_creacion)
		if e != nil {
			ctx.JSON(400, gin.H{
				"error": e,
			})
			return
		}
		ctx.JSON(200, product)
	}

}

func (p *Producto) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, _ := strconv.ParseInt(id, 10, 64)
		product, e := p.service.GetById(int(intId))
		if e != nil {
			ctx.JSON(400, gin.H{
				"error": e,
			})
			return
		}
		ctx.JSON(200, product)
	}

}
