package handler

import (
	transacciones "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transacciones"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaCreacion     string  `json:"fecha_creacion"`
}

type Transaccion struct {
	service transacciones.Service
}

func NewTransaccion(t transacciones.Service) *Transaccion {
	return &Transaccion{
		service: t,
	}
}

func (t *Transaccion) Load() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "" {
			if token == "39470939" {
				t, err := t.service.Load()

				if err != nil {
					c.JSON(404, gin.H{
						"error": err.Error(),
					})
					return
				}
				c.JSON(200, t)

			} else {
				c.JSON(401, gin.H{
					"error": "token incorrecto",
				})
			}

		} else {
			c.JSON(401, gin.H{
				"error": "debes ingresas token en el header",
			})
		}

	}
}

func (t *Transaccion) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "" {
			if token == "39470939" {

				t, err := t.service.GetAll()

				if err != nil {
					c.JSON(404, gin.H{
						"error": err.Error(),
					})
					return
				}
				c.JSON(200, t)

			} else {
				c.JSON(401, gin.H{
					"error": "token incorrecto",
				})
			}

		} else {
			c.JSON(401, gin.H{
				"error": "debes ingresas token en el header",
			})
		}

	}
}

func (t *Transaccion) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		token := c.GetHeader("token")

		if token != "" {
			if token == "39470939" {

				if err := c.Bind(&req); err != nil {
					c.JSON(404, gin.H{
						"error": err.Error(),
					})
					return
				} else {
					switch {
					case req.Monto == 0.0:
						c.JSON(401, gin.H{
							"error": "no se puede poner el monto vacio",
						})
					case req.Emisor == "":
						c.JSON(401, gin.H{
							"error": "no se puede emitir al emisor",
						})
					case req.Moneda == "":
						c.JSON(401, gin.H{
							"error": "no se puede emitir el tipo de moneda",
						})
					case req.Receptor == "":
						c.JSON(401, gin.H{
							"error": "no se puede emitir al receptor",
						})
					default:
						t, err := t.service.Store(req.ID, req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaCreacion)
						if err != nil {
							c.JSON(404, gin.H{
								"error": err.Error(),
							})
							return
						}
						c.JSON(200, t)

					}

				}

			} else {
				c.JSON(401, gin.H{
					"error": "token incorrecto",
				})
			}

		} else {
			c.JSON(401, gin.H{
				"error": "debes ingresas token en el header",
			})
		}

	}
}
