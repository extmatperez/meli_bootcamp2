package handler

import (
	"strconv"

	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/8_goweb3/ClaseTM/ProyectoEstructura/internal/producto"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fecha_creacion"`
}

type ProductoController struct {
	service producto.ServiceProducto
}

func NewProductoController(ser producto.ServiceProducto) *ProductoController {
	return &ProductoController{service: ser}
}

func (prodController *ProductoController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		personas, err := prodController.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, personas)
		}
	}
}

func (prodController *ProductoController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var produc request

		err := ctx.ShouldBindJSON(&produc)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := prodController.service.Store(produc.Nombre, produc.Color, produc.Precio, produc.Stock, produc.Codigo, produc.Publicado, produc.FechaCreacion)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (prodController *ProductoController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var produc request

		id, err := strconv.Atoi(ctx.Param("id"))

		if err == nil {
			err := ctx.ShouldBindJSON(&produc)

			if err == nil {
				prod, err := prodController.service.Update(id, produc.Nombre, produc.Color, produc.Precio, produc.Stock, produc.Codigo, produc.Publicado, produc.FechaCreacion)

				if err == nil {
					ctx.JSON(200, prod)
				} else {
					ctx.JSON(404, gin.H{
						"message": err.Error(),
					})
				}
			} else {
				ctx.String(400, "Error con los datos del body")
			}
		} else {
			ctx.String(400, "Error con el id")
		}

	}
}

func (prodController *ProductoController) Detele() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err == nil {
			err := prodController.service.Delete(id)
			if err == nil {
				ctx.String(200, "Producto %d borrado exitosamente", id)
			} else {
				ctx.JSON(400, gin.H{
					"message": err.Error(),
				})
			}
		} else {
			ctx.String(400, "Error con el id")
		}
	}
}

func (prodController *ProductoController) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		var prod request
		errorCast := ctx.ShouldBindJSON(&prod)
		if err == nil {
			if errorCast == nil {
				productoMod, errorMod := prodController.service.UpdateNameAndPrice(id, prod.Nombre, prod.Precio)
				if errorMod == nil {
					ctx.JSON(200, gin.H{
						"message":  "Campos Modificados correctamente",
						"producto": productoMod,
					})
				} else {
					ctx.JSON(200, gin.H{
						"message": errorMod.Error(),
					})
				}
			} else {
				ctx.String(200, "Hubo un error con body, asegurese de pasar bien los campos nombre y precio")
			}

		} else {
			ctx.String(200, "Hubo un error con el id")
		}
	}
}
