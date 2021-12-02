package handler

import (
	"fmt"
	"os"
	"strconv"

	product "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTM/internal/products"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTM/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}
type Product struct {
	service product.Service
}

func NewProduct(ser product.Service) *Product {
	return &Product{service: ser}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(400, web.NewReponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
		} else {
			ctx.JSON(200, web.NewReponse(200, products, ""))
		}
	}
}
func (controller *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		var prod request
		err := ctx.ShouldBindJSON(&prod)

		if err != nil {
			ctx.JSON(400, web.NewReponse(400, nil, fmt.Sprintf("hubo un error al querer guardar la persona %v", err)))
		} else {
			response, err := controller.service.Store(prod.Nombre, prod.Color, prod.Precio, prod.Stock, prod.Codigo, prod.Publicado, prod.FechaCreacion)
			if err != nil {
				ctx.JSON(400, web.NewReponse(400, nil, fmt.Sprintf("No se pudo cargar la persona %v", err)))
			} else {
				ctx.JSON(200, web.NewReponse(200, response, ""))
			}
		}
	}
}

/*
func (controller *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pro request
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.String(400, "Error el id es invalido")
		}
		err = ctx.ShouldBindJSON(&pro)
		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			productUpdate, err := controller.service.Update(int(id), pro.Nombre, pro.Color, pro.Precio, pro.Stock, pro.Codigo, pro.Publicado, pro.FechaCreacion)
			if err != nil {
				ctx.String(400, "error:%v", err)
			} else {
				ctx.JSON(200, productUpdate)

			}
		}
	}
}*/

func (controller *Product) UpdateNombre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		var pro request
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewReponse(404, nil, "Error el id es invalido"))
		}
		err = ctx.ShouldBindJSON(&pro)
		if err != nil {
			ctx.JSON(400, web.NewReponse(400, nil, "Error el body"))
		} else {
			if pro.Nombre == "" {
				ctx.JSON(400, web.NewReponse(400, nil, "Error el nombre no puede estar vacio"))
				return
			}
			productUpdateNombre, err := controller.service.UpdateNombre(int(id), pro.Nombre)
			if err != nil {
				ctx.JSON(404, web.NewReponse(404, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewReponse(200, productUpdateNombre, ""))

			}
		}
	}
}
func (controller *Product) UpdatePrecio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		var pro request
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewReponse(404, nil, "Error el id es invalido"))
			//ctx.String(400, "Error el id es invalido")
		}
		err = ctx.ShouldBindJSON(&pro)
		if err != nil {
			ctx.JSON(400, web.NewReponse(400, nil, "Error el body"))
			//ctx.String(404, "Error en el body")
		} else {
			if pro.Precio == 0 {
				ctx.JSON(400, web.NewReponse(400, nil, "Error el precio no puede ser cero"))
				//ctx.String(400, "Error el precio no puede ser cero")
				return
			}
			productUpdatePrecio, err := controller.service.UpdatePrecio(int(id), pro.Precio)
			if err != nil {
				ctx.JSON(404, web.NewReponse(404, nil, err.Error()))
				//ctx.String(404, "error:%v", err)
			} else {
				ctx.JSON(200, web.NewReponse(200, productUpdatePrecio, ""))
				//ctx.String(404, "error:%v", err)
			}
		}
	}
}
func (controller *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		var pro request
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewReponse(400, nil, "Error el id es invalido"))
		}
		err = ctx.Bind(&pro)
		if err != nil {
			ctx.JSON(404, web.NewReponse(404, nil, err.Error()))
			return
		}
		if pro.Nombre == "" {
			ctx.JSON(404, web.NewReponse(404, nil, "el nombre no debe estar vacio"))
			return
		}
		if pro.Color == "" {
			ctx.JSON(404, web.NewReponse(404, nil, "el color no debe estar vacio"))
			return
		}
		if pro.Precio == 0 {
			ctx.JSON(404, web.NewReponse(404, nil, "el precio no debe ser cero"))
			return
		}
		if pro.Stock == "" {
			ctx.JSON(404, web.NewReponse(404, nil, "el stock no debe estar vacio"))
			return
		}
		if pro.Codigo == "" {
			ctx.JSON(404, web.NewReponse(404, nil, "el codigo no debe estar vacio"))
			return
		}
		if pro.FechaCreacion == "" {
			ctx.JSON(404, web.NewReponse(404, nil, "la fecha no debe estar vacia"))
			return
		}
		productUpdate, err := controller.service.Update(int(id), pro.Nombre, pro.Color, pro.Precio, pro.Stock, pro.Codigo, pro.Publicado, pro.FechaCreacion)

		if err != nil {
			ctx.JSON(404, web.NewReponse(404, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewReponse(200, productUpdate, ""))
		}
	}
}
func (controller *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewReponse(400, nil, "Error el id es invalido"))
		}
		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewReponse(400, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewReponse(400, nil, fmt.Sprintf("La persona %d ha sido eliminada", id)))
		}
	}
}

// funciones generales
func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(400, web.NewReponse(400, nil, "falta token"))
		//ctx.String(400, "Fatal token")
		return false
	}
	tokenEnv := os.Getenv("TOKEN")
	if token != tokenEnv {
		ctx.JSON(400, web.NewReponse(404, nil, "token incorrecto"))
		//ctx.String(400, "Token incorrecto")
		return false
	}
	return true
}
