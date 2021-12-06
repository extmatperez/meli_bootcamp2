package handler

import (
	"os"
	"strconv"

	productos "github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/go-web/clase3/goModularizadoEnCapas/Internal/productos"
	"github.com/gin-gonic/gin"
)

type product struct {
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

type Producto struct {
	service productos.Service
}

func NewPersona(ser productos.Service) *Producto {
	return &Producto{service: ser}
}

func ValidarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, "El token no es correcto")
		return false
	}
	return true
}

func (per *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ValidarToken(ctx) {
			ctx.JSON(401, "")
			return
		}
		personas, err := per.service.GetAll()
		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, personas)
		}
	}
}

func (controller *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ValidarToken(ctx) {
			ctx.JSON(401, "")
			return
		}
		var producto productos.Product
		err := ctx.ShouldBindJSON(&producto)
		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			if producto.Nombre == "" {
				ctx.JSON(400, gin.H{"error": "No ingresó el nombre"})
				return
			}
			if producto.Stock == "0" {
				ctx.JSON(400, gin.H{"error": "El stock no puede ser cero."})
				return
			}
			response, err := controller.service.Store(producto)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (controller *Producto) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ValidarToken(ctx) {
			ctx.JSON(401, "")
			return
		}
		var productoAux productos.Product
		err := ctx.ShouldBindJSON(&productoAux)
		varID, _ := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Update(varID, productoAux)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (controller *Producto) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ValidarToken(ctx) {
			ctx.JSON(401, "")
			return
		}
		var nameUpdate = ctx.GetHeader("nameUpdate")
		varID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(400, "No se pudo convertir ID a integer %v", err)
		} else {
			err := controller.service.UpdateName(varID, nameUpdate)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.String(200, "Se actualizo el producto con el id %v", varID)
			}
		}

	}
}

func (controller *Producto) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ValidarToken(ctx) {
			ctx.JSON(401, "")
			return
		}
		varID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(400, "No se pudo convertir ID a integer %v", err)
		} else {
			err := controller.service.Delete(varID)
			if err != nil {
				ctx.String(400, "No se ha encontraro un producto para eliminar con id %v", varID)
			} else {
				ctx.String(200, "Se elimino el producto con el id %v", varID)
			}
		}

	}
}
