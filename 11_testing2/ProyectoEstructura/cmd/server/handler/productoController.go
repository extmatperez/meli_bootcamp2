package handler

import (
	"fmt"
	"os"
	"strconv"

	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/11_testing2/ProyectoEstructura/internal/producto"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/11_testing2/ProyectoEstructura/pkg/web"
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

func validarToken(ctx *gin.Context) (bool, int, error) {

	tok := ctx.GetHeader("token")
	if os.Getenv("TOKEN") == tok {
		return true, 200, nil
	} else {
		return false, 403, fmt.Errorf("Token incorrecto: %s", tok)
	}
}

func NewProductoController(ser producto.ServiceProducto) *ProductoController {
	return &ProductoController{service: ser}
}

// GetAll godoc
// @Summary List products
// @Tags Products
// @Description Get all products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /api/productos [get]
func (prodController *ProductoController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		validado, status, err := validarToken(ctx)

		if validado {
			personas, err := prodController.service.GetAll()

			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			} else {
				ctx.JSON(200, web.NewResponse(200, personas, ""))
			}
		} else {
			ctx.JSON(status, web.NewResponse(status, nil, fmt.Sprintf("Hubo un error %v", err)))
		}
	}
}

// Store godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (prodController *ProductoController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		validado, status, err := validarToken(ctx)
		if validado {
			var produc request

			err := ctx.ShouldBindJSON(&produc)

			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar una persona %v", err)))
			} else {
				response, err := prodController.service.Store(produc.Nombre, produc.Color, produc.Precio, produc.Stock, produc.Codigo, produc.Publicado, produc.FechaCreacion)
				if err != nil {
					ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar la persona %v", err)))
				} else {
					ctx.JSON(200, web.NewResponse(200, response, ""))
				}
			}
		} else {
			ctx.JSON(status, web.NewResponse(status, nil, fmt.Sprintf("Hubo un error de autorizacion:  %v", err.Error())))
		}
	}
}

// Store godoc
// @Summary Update product
// @Tags Products
// @Description Update product
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Router /products/{id} [put]
func (prodController *ProductoController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		validado, status, err := validarToken(ctx)

		if validado {
			var produc request

			id, err := strconv.Atoi(ctx.Param("id"))

			if err == nil {
				err := ctx.ShouldBindJSON(&produc)

				if err == nil {
					prod, err := prodController.service.Update(id, produc.Nombre, produc.Color, produc.Precio, produc.Stock, produc.Codigo, produc.Publicado, produc.FechaCreacion)

					if err == nil {
						ctx.JSON(200, web.NewResponse(200, prod, ""))
					} else {
						ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("Hubo un problema al updatear  %v", err.Error())))

					}
				} else {
					ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error con los datos del body  %v", err.Error())))
				}
			} else {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error con el id:  %v", err.Error())))
			}
		} else {
			ctx.JSON(status, web.NewResponse(status, nil, fmt.Sprintf("Hubo un error de autorizacion:  %v", err.Error())))
		}

	}
}

func (prodController *ProductoController) Detele() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		validado, status, err := validarToken(ctx)

		if validado {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err == nil {
				err := prodController.service.Delete(id)
				if err == nil {
					ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("Producto %d borrado exitosamente", id), ""))

				} else {
					ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al borrar:  %v", err.Error())))
				}
			} else {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error con el id:  %v", err.Error())))
			}
		} else {
			ctx.JSON(status, web.NewResponse(status, nil, fmt.Sprintf("Hubo un error de autorizacion:  %v", err.Error())))
		}

	}
}

func (prodController *ProductoController) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validado, status, err := validarToken(ctx)

		if validado {
			id, err := strconv.Atoi(ctx.Param("id"))
			var prod request
			errorCast := ctx.ShouldBindJSON(&prod)
			if err == nil {
				if errorCast == nil {
					productoMod, errorMod := prodController.service.UpdateNameAndPrice(id, prod.Nombre, prod.Precio)
					if errorMod == nil {
						ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("campos modificados correctamente", productoMod), ""))

					} else {
						ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error con el id:  %v", err.Error())))
					}
				} else {
					ctx.String(400, "Hubo un error con body, asegurese de pasar bien los campos nombre y precio")
				}

			} else {
				ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("Hubo un error con el id:  %v", err.Error())))
			}
		} else {
			ctx.JSON(status, web.NewResponse(status, nil, fmt.Sprintf("Hubo un error de autorizacion:  %v", err.Error())))
		}

	}
}
