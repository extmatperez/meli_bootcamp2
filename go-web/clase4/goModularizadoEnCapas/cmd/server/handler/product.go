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
	//	err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("error al intentar cargar el archivo .env")
	//		return false
	//	}
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
		}
		var producto productos.Product
		err := ctx.ShouldBindJSON(&producto)
		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
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
