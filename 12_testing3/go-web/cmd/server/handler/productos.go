package handler

import (
	"os"
	"strconv"

	productos "github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/12_testing3/go-web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/pkg/web"
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

func validateToken(tokenReq string) bool {
	tokenEnv := os.Getenv("TOKEN")
	return tokenEnv == tokenReq
}

// Store godoc
// @Summary Stores products
// @Tags Productos
// @Description posts product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenReq := ctx.Request.Header.Get("token")
		if !validateToken(tokenReq) {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Productos no encontrados"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// Store godoc
// @Summary Stores products
// @Tags Productos
// @Description posts product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos/{id} [post]
func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenReq := ctx.Request.Header.Get("token")
		if !validateToken(tokenReq) {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		p, err := c.service.Store(req.ID, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Creado)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, p)
	}
}

// Edit godoc
// @Summary Edits a product
// @Tags Productos
// @Description posts product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos/{id} [put]
func (c *Product) Edit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenReq := ctx.Request.Header.Get("token")
		if !validateToken(tokenReq) {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}
		var req request
		id, existeId := ctx.GetQuery("id")
		if !existeId {
			ctx.JSON(400, web.NewResponse(400, nil, "Producto no especificado"))
			return
		}
		parsedId, _ := strconv.Atoi(id)
		err := ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		p, err := c.service.Edit(parsedId, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.Creado)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, p)
	}
}

// Delete godoc
// @Summary Deletes a product
// @Tags Productos
// @Description Deletes product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos/{id} [delete]
func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenReq := ctx.Request.Header.Get("token")
		if !validateToken(tokenReq) {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}
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

// Delete godoc
// @Summary Edits prudct name & price
// @Tags Productos
// @Description posts product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos/{id} [patch]
func (c *Product) Change() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenReq := ctx.Request.Header.Get("token")
		if !validateToken(tokenReq) {
			ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
			return
		}
		var req request
		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "No se ha especificado un prodcuto"))
			return
		}
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, web.NewResponse(400, nil, err.Error()))
			return
		}
		err = ctx.ShouldBind(&req)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		cambios, err := c.service.Change(parsedId, req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, cambios)
	}
}
