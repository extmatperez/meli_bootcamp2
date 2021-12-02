package handler

import (
	"strconv"

	product "github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/internal/producto"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service product.Service
}

type Request struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

func NewProduct(ser product.Service) *Product {
	return &Product{service: ser}
}

func (pro *Product) LoadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := pro.service.LoadFile()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, "El archivo fue cargado con exito")
		}
	}
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

func (pro *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Request

		err := ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar los datos %v", err)
		} else {
			response, err := pro.service.Store(product.Name, product.Color, product.Price, product.Stock, product.Code, product.IsPublished, product.CreatedAt)
			if err != nil {
				ctx.String(400, "No se pudo guardad el producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

func (pro *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idStr := ctx.Param("ID")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(400, "Hubo un error")
			return
		}

		mes, err := pro.service.Delete(int64(id))

		if err != nil {
			ctx.String(400, "No se pudo guardad el producto %v", err)
		} else {
			ctx.JSON(200, mes)
		}
	}
}

func (pro *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Request
		idStr := ctx.Param("ID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(400, "Hubo un error")
			return
		}

		err = ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar los datos %v", err)
		} else {
			response, err := pro.service.Update(int64(id), product.Name, product.Color, product.Price, product.Stock, product.Code, product.IsPublished, product.CreatedAt)
			if err != nil {
				ctx.String(400, "No se pudo actualizar el producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

func (pro *Product) UpdateNombre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Request
		idStr := ctx.Param("ID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(400, "Hubo un error")
			return
		}

		err = ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar los datos %v", err)
		} else {
			response, err := pro.service.UpdateNombre(int64(id), product.Name)
			if err != nil {
				ctx.String(400, "No se pudo actualizar el nombre del producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}
