package handler

import (
	products "github.com/extmatperez/meli_bootcamp2/tree/parra_diego/7_goweb2/Tarde/ejercicio_2/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	// Id      int     `json:"id"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Code    string  `json:"code"`
	Publish bool    `json:"publish"`
	Date    string  `json:"date"`
}
type Product struct {
	service products.Service
}

func NewProduct(ser products.Service) *Product {
	return &Product{service: ser}
}

func (pro *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := pro.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, products)
		}
	}
}

func (controller *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var produ request

		err := ctx.ShouldBindJSON(&produ)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar un producto %v", err)
		} else {
			response, err := controller.service.Store(produ.Name, produ.Color, produ.Price, produ.Stock, produ.Code, produ.Publish, produ.Date)
			if err != nil {
				ctx.String(400, "No se pudo cargar un producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}
