package handler

import (
	"fmt"
	"os"
	"strconv"

	internal "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/7_goweb2/ejTTmodified/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/7_goweb2/ejTTmodified/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre string  `json:"nombre"`
	Color  string  `json:"color" `
	Precio float64 `json:"precio" `
}
type Producto struct {
	service internal.Service
}

func NewProducto(serv internal.Service) *Producto {
	return &Producto{service: serv}
}

func (p *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, products, ""))
	}

}

func (p *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		var prod request

		err := ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		product, e := p.service.Store(prod.Nombre, prod.Color, prod.Precio)
		if e != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, product, ""))
	}

}

func (p *Producto) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		id := ctx.Param("id")
		intId, _ := strconv.ParseInt(id, 10, 64)
		product, e := p.service.GetById(int(intId))
		if e != nil {
			ctx.JSON(400, web.NewResponse(400, nil, e.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, product, ""))
	}

}
func (p *Producto) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		var prod request
		err := ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		id := ctx.Param("id")
		intId, _ := strconv.ParseInt(id, 10, 64)
		product, e := p.service.Update(int(intId), prod.Nombre, prod.Color, prod.Precio)
		if e != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, product, ""))
	}

}

func (p *Producto) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		id := ctx.Param("id")
		intId, _ := strconv.ParseInt(id, 10, 64)
		err := p.service.Delete(int(intId))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("se ha borrado exitosamente el registro de id %v", id), ""))
	}

}
func (p *Producto) UpdateNombrePrecio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		var prod request
		id, _ := strconv.Atoi(ctx.Param("id"))
		err := ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if prod.Nombre == "" || prod.Precio == 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		productoActualizado, e := p.service.UpdateNombrePrecio(id, prod.Nombre, prod.Precio)
		if e != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, productoActualizado, ""))

	}
}

func validateToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")

	if token == "" {
		ctx.JSON(400, web.NewResponse(401, nil, "Token not entered"))
		return false
	}
	if token != os.Getenv("TOKEN") {
		ctx.JSON(400, web.NewResponse(401, nil, "Not valid token"))
		return false
	}
	return true

}
