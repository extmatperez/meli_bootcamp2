package handler

import (
	product "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/7_goweb2/go-web-TT/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}
type Product struct {
	service product.Service
}

func NewProduct(ser product.Service) *Product {
	return &Product{service: ser}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll()
		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, products)
		}
	}
}
func (controller *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod request
		err := ctx.ShouldBindJSON(&prod)

		if err != nil {
			ctx.String(400, "hubo un error al querer guardar la persona %v", err)
		} else {
			response, err := controller.service.Store(prod.Nombre, prod.Color, prod.Precio, prod.Stock, prod.Codigo, prod.Publicado, prod.FechaCreacion)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}
