package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	products "github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/internal/products"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Nombre        string `json:"nombre" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Precio        int    `json:"precio" binding:"required"`
	Stock         int    `json:"stock" binding:"required"`
	Codigo        string `json:"codigo" binding:"required"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion" binding:"required"`
}

type RequestPatchNamePrice struct {
	Nombre string `json:"nombre" binding:"required"`
	Precio int    `json:"precio" binding:"required"`
}

type Product struct {
	serv    products.Service
	servSQL products.ServiceSQL
}

func NewProduct(s products.Service) *Product {
	return &Product{serv: s}
}

func NewProductSQL(s products.ServiceSQL) *Product {
	return &Product{servSQL: s}
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
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		prods, err := p.serv.GetAll()

		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prods, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body Request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newRequest Request
		err := c.ShouldBindJSON(&newRequest)
		if err != nil {
			validRequired(err.Error(), newRequest, c)
		} else {

			prod, err := p.servSQL.Store(newRequest.Nombre, newRequest.Color, newRequest.Precio, newRequest.Stock, newRequest.Codigo, newRequest.Publicado, newRequest.FechaCreacion)

			if err != nil {
				c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			} else {
				c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))
			}

		}

	}
}

// Get by id Product godoc
// @Summary get product by id
// @Tags Products
// @Description this handler get a product by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /products/{id} [get]
func (p *Product) FindById() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id invalido"))
			return
		}

		prod, err := p.serv.FindById(id)

		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))
		}

	}
}

// Update Product godoc
// @Summary Update product
// @Tags Products
// @Description this handler updates a product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body Request true "Product to edit"
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /products/{id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updateRequest Request
		err := c.ShouldBindJSON(&updateRequest)
		if err != nil {
			validRequired(err.Error(), updateRequest, c)
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id invalido"))
			return
		}

		prod, err := p.serv.Update(id, updateRequest.Nombre, updateRequest.Color, updateRequest.Precio, updateRequest.Stock, updateRequest.Codigo, updateRequest.Publicado, updateRequest.FechaCreacion)

		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id invalido"))
			return
		}

		err = p.serv.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, "Eliminado exitosamente", ""))

	}
}

func (p *Product) Filter() gin.HandlerFunc {
	return func(c *gin.Context) {

		prods, err := p.serv.Filter(c.Request.URL.Query())
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prods, ""))
	}
}

func (p *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id invalido"))

			return
		}

		var updateRequest RequestPatchNamePrice
		err = c.ShouldBindJSON(&updateRequest)
		if err != nil {
			validRequired(err.Error(), updateRequest, c)
			return
		}

		prod, err := p.serv.UpdateNameAndPrice(id, updateRequest.Nombre, updateRequest.Precio)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))

	}
}

func validRequired(err string, Request interface{}, c *gin.Context) {

	if strings.Contains(err, "required") {
		tipos := reflect.TypeOf(Request)
		i := 0
		errores := ""
		for i = 0; i < tipos.NumField(); i++ {
			if strings.Contains(err, tipos.Field(i).Name) {
				errores = errores + fmt.Sprintf("El campo %s es requerido. ", tipos.Field(i).Name)
			}
		}
		c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, errores))

	} else {
		c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err))
	}
}
