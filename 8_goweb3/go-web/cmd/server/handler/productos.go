package handler

import (
	"strconv"

	productos "github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/go-web/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Color     string `json:"color"`
	Precio    string `json:"precio"`
	Stock     int    `json:"stock"`
	Codigo    string `json:"codigo"`
	Publicado bool   `json:"publicado"`
	Creado    string `json:"creado"`
}

type Product struct {
	service productos.Service
}

func NewProduct(p productos.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.ID, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Creado)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Edit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		id, existeId := ctx.GetQuery("id")
		if !existeId {
			ctx.JSON(400, gin.H{"error": "Especifique ID"})
			return
		}
		parsedId, _ := strconv.Atoi(id)
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p, err := c.service.Edit(parsedId, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Creado)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(400, gin.H{"error": "No se ha seleccionado un producto"})
			return
		}
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = c.service.Delete(parsedId)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, nil)
	}
}
