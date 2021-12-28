package handler

import (
	"fmt"
	"os"
	"strconv"

	producto "github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre            string  `json:"nombre"`
	Color             string  `json:"color"`
	Precio            float64 `json:"precio"`
	Stock             int     `json:"stock"`
	Codigo            string  `json:"codigo"`
	Publicado         bool    `json:"publicado"`
	Fecha_de_creacion string  `json:"fecha_de_creacion"`
}

type Producto struct {
	service producto.Service
}

func NewProducto(ser producto.Service) *Producto {
	return &Producto{service: ser}
}

func validarToken(c *gin.Context) bool {
	token := c.GetHeader("token")

	if token == "" {
		c.JSON(400, web.NewResponse(400, nil, "No envio el token"))
		// c.String(400, "Falta token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if tokenENV != token {
		c.JSON(404, web.NewResponse(404, nil, "Token incorrecto"))
		// c.String(400, "Token incorrecto")
		return false
	}
	return true
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (prod *Producto) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}

		prod, err := prod.service.GetAll()

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			// c.String(400, "Hubo un error %v", err)
		} else {
			c.JSON(200, web.NewResponse(200, prod, ""))
		}
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /addProductos [post]
func (controller *Producto) Store() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}
		var prod request
		err := c.ShouldBindJSON(&prod)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar un producto %v", err)))
			// c.String(400, "Hubo un error al querer cargar un producto %v", err)
		} else {
			response, err := controller.service.Store(prod.Stock, prod.Nombre, prod.Codigo, prod.Color, prod.Fecha_de_creacion, prod.Precio, prod.Publicado)

			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar el producto: %v", err)))
				// c.String(400, "No se pudo cargar el producto: %v", err)
			} else {
				c.JSON(200, web.NewResponse(200, response, ""))
			}
		}
	}
}

// ModifyProducts godoc
// @Summary Modify products
// @Tags Products
// @Description modify products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to Modify"
// @Success 200 {object} web.Response
// @Router /modify/:id [put]
func (controller *Producto) Modify() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}
		var prod request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "Hubo un error, el id es invalido"))
			// c.String(404, "Hubo un error, el id es invalido")
		}

		err = c.ShouldBindJSON(&prod)

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "Hubo un error en el body"))
			// c.String(400, "Hubo un error en el body")
		} else {
			prodActualizado, err := controller.service.Modify(int(id), prod.Stock, prod.Nombre, prod.Codigo, prod.Color, prod.Fecha_de_creacion, prod.Precio, prod.Publicado)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, "No se pudo modificar el producto"))
				// c.JSON(400, err.Error())
			} else {
				c.JSON(200, web.NewResponse(200, prodActualizado, ""))
			}
		}
	}
}

// Modify Name Price Products godoc
// @Summary Modify Name Price products
// @Tags Products
// @Description Modify Name Price products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to Modify Name Price"
// @Success 200 {object} web.Response
// @Router /modifyNaPr/:id [patch]
func (controller *Producto) ModifyNamePrice() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}
		var prod request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "Hubo un error, el id es invalido"))
			// c.String(404, "Hubo un error, el id es invalido")
		}

		err = c.ShouldBindJSON(&prod)

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "Hubo un error en el body"))
			// c.String(400, "Hubo un error en el body")
		} else {

			prodActualizado, err := controller.service.ModifyNamePrice(int(id), prod.Nombre, prod.Precio)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, "No se pudo modificar el producto"))
				// c.JSON(400, err.Error())
			} else {
				c.JSON(200, web.NewResponse(200, prodActualizado, ""))
			}
		}
	}
}

// Delete Products godoc
// @Summary Delete products
// @Tags Products
// @Description Delete products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to Delete"
// @Success 200 {object} web.Response
// @Router /delete/:id [delete]
func (controller *Producto) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, "Hubo un error, el id es invalido"))
			// c.String(404, "Hubo un error, el id es invalido")
		}

		err = controller.service.Delete(int(id))
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "No se pudo eliminar el producto"))
			// c.JSON(400, err.Error())
		} else {
			c.JSON(200, web.NewResponse(200, nil, "El producto ha sido eliminado"))
		}
	}
}
