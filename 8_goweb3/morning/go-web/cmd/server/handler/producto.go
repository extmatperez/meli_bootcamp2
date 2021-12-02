package handler

import (
	"os"
	"strconv"

	productos "github.com/extmatperez/meli_bootcamp2/7_goweb2/afternoon/go-web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/afternoon/go-web/pkg/web"
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

func ValidaToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		/* ctx.JSON(401, gin.H{
			"error": "token inválido",
		}) */
		return false
	}
	return true
}

func (prod *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}
		response, err := prod.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			//ctx.JSON(400, err)
			return
		}
		ctx.JSON(200, web.NewResponse(200, response, ""))
		//ctx.JSON(200, response)
	}
}

func (prod *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}
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

		if !ValidaToken(ctx) {
			return
		}
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

		if !ValidaToken(ctx) {
			return
		}
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

		if !ValidaToken(ctx) {
			return
		}
		myId, err := strconv.Atoi(ctx.Param("id"))

		response, err := prod.service.Delete(myId)

		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		ctx.JSON(200, response)
	}
}
