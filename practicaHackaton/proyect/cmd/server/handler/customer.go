package handler

import (
	"strconv"

	customers "github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/customers"
	models "github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/models"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/web"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	service customers.Service
}

func NewCustomer(ser customers.Service) *Customer {
	return &Customer{service: ser}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Customer) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}
		response, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, response, ""))
	}
}

func (c *Customer) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}
		myId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID type"))
			return
		}

		response, err := c.service.GetByID(myId)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if response.ID == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Not Found"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, response, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Customer) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}
		var req models.RequestCustomer
		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		response, err := c.service.Store(req.First_Name, req.Last_Name, req.Condition)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, response, ""))
	}
}

func (c *Customer) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}
		var req models.RequestCustomer
		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		myId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID type"))
			return
		}

		response, err := c.service.Update(ctx, myId, req.First_Name, req.Last_Name, req.Condition)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, response, ""))
	}
}

func (c *Customer) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}
		myId, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID: "+err.Error()))
			return
		}

		responseGet, err := c.service.GetByID(myId)
		if responseGet.ID == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "Not Found"))
			return
		}

		err = c.service.Delete(ctx, myId)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, "", ""))
	}
}
