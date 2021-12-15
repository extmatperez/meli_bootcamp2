package handler

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/products"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fecha_de_creacion"`
}

type Product struct {
	service products.Service
}

func NewProduct(prod products.Service) *Product {
	return &Product{
		service: prod,
	}
}

func validation(req request) string {
	reqValue := reflect.ValueOf(req)
	for i := 0; i < reqValue.NumField(); i++ {
		value := reqValue.Field(i).Interface()
		tipe := reflect.TypeOf(value).Kind()
		if fmt.Sprint(tipe) == "string" {
			if value == "" {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "int" {
			if value.(int) < 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "float64" {
			if value.(float64) == 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "boolean" {
			if value.(bool) == false {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		}
	}
	return ""
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
func (prod *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, fmt.Sprintf("Token invalido")))
			return
		}
		p, err := prod.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(400, nil, fmt.Sprintf("error %v", err)))
			return
		}
		c.JSON(200, web.NewResponse(200, p, ""))
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
func (controller *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, fmt.Sprintf("Token invalido")))
			return
		}
		var prod request
		prod.Stock = -1
		err := c.ShouldBindJSON(&prod)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar el producto %v", err)))
		} else {
			validated := validation(prod)
			if validated != "" {
				c.JSON(400, web.NewResponse(400, nil, validated))
				return
			}
			response, err := controller.service.Store(prod.Nombre, prod.Color, prod.Precio, prod.Stock, prod.Codigo, prod.Publicado, prod.FechaDeCreacion)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar el producto %v", err)))
			} else {
				c.JSON(200, web.NewResponse(200, response, ""))
			}
		}
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
// @Router /products [put]
func (controller *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(400, nil, fmt.Sprintf("Token invalido")))
			return
		}
		var prod request
		id, err1 := strconv.Atoi(c.Param("id"))
		if err1 != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo obtener el id %v", err1)))
		} else {
			err2 := c.ShouldBindJSON(&prod)
			validated := validation(prod)
			if validated != "" {
				c.JSON(400, web.NewResponse(400, nil, validated))
				return
			}
			if err2 != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar el producto %v", err2)))
			} else {
				fmt.Println("PRODUCTO", prod)
				response, err := controller.service.Update(id, prod.Nombre, prod.Color, prod.Precio, prod.Stock, prod.Codigo, prod.Publicado, prod.FechaDeCreacion)
				if err != nil {
					c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar el producto %v", err)))

				} else {
					c.JSON(200, web.NewResponse(200, response, ""))
				}
			}
		}
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
// @Router /products [patch]
func (controller *Product) UpdateProd() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, fmt.Sprintf("Token invalido")))
			return
		}
		var prod request
		id, err1 := strconv.Atoi(c.Param("id"))
		if err1 != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo obtener el id %v", err1)))
		} else {
			err2 := c.ShouldBindJSON(&prod)
			if err2 != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar el producto %v", err2)))
			} else {
				response, err := controller.service.UpdateProd(id, prod.Nombre, prod.Precio)
				if err != nil {
					c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo actualizar el producto %v", err)))
				} else {
					c.JSON(200, web.NewResponse(200, response, ""))
				}
			}
		}
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
// @Router /products [delete]
func (controller *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(401, web.NewResponse(401, nil, fmt.Sprintf("Token invalido")))
			return
		}
		id, err1 := strconv.Atoi(c.Param("id"))
		if err1 != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo obtener el id %v", err1)))
		} else {
			err := controller.service.Delete(id)
			if err != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo eliminar el producto %v", err)))
			} else {
				c.JSON(200, web.NewResponse(200, "Producto eliminado correctamente", ""))
			}
		}
	}
}
