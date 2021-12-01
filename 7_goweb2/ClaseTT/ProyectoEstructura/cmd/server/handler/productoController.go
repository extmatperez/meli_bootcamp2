package handler

import (
	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/7_goweb2/ClaseTT/ProyectoEstructura/internal/producto"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

type ProductoController struct {
	service producto.ServiceProducto
}

func NewProductoController(ser producto.ServiceProducto) *ProductoController {
	return &ProductoController{service: ser}
}

func (prodController *ProductoController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		personas, err := prodController.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, personas)
		}
	}
}

func (prodController *ProductoController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var produc request

		err := ctx.ShouldBindJSON(&produc)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := prodController.service.Store(produc.Nombre, produc.Color, produc.Precio, produc.Stock, produc.Codigo, produc.Publicado, produc.FechaCreacion)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}
