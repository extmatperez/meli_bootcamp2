package handler

import (
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/tree/parra_diego/8_goweb3/TM/ejercicio_1/internal/productos"
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

func (controller *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var pro request
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(404, "Id incorrecto")
		}
		err = ctx.ShouldBindJSON(&pro)

		if pro.Name == "" || pro.Color == "" || pro.Price == 0 || pro.Stock == 0 || pro.Code == "" || pro.Date == "" {
			ctx.String(400, "Todos los campos son requeridos")
		} else {

			if err != nil {
				ctx.String(400, "Error en el body")
			} else {
				produUpdate, err := controller.service.Update(int(id), pro.Name, pro.Color, pro.Price, pro.Stock, pro.Code, pro.Publish, pro.Date)
				if err != nil {
					ctx.JSON(400, err.Error())
				} else {
					ctx.JSON(200, produUpdate)
				}
			}
		}
	}
}
