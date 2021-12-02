package handler

import (
	"strconv"

	transacciones "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/internal/transacciones"
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
							"error": "no se puede omitir el tipo de moneda",
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

func (t *Transaccion) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		token := c.GetHeader("token")

		if token != "" {
			if token == "39470939" {

				id, err := strconv.ParseInt(c.Param("id"), 10, 64)

				if err != nil {
					c.JSON(400, "el id es invalido")
				} else {
					if err := c.ShouldBindJSON(&req); err != nil {
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
								"error": "no se puede omitir al emisor",
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
							t, err := t.service.Update(int(id), req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaCreacion)
							if err != nil {
								c.JSON(404, gin.H{
									"error": err.Error(),
								})
								return
							}
							c.JSON(200, t)

						}

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

func (t *Transaccion) UpdateCod() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		token := c.GetHeader("token")

		if token != "" {
			if token == "39470939" {

				id, err := strconv.ParseInt(c.Param("id"), 10, 64)

				if err != nil {
					c.JSON(400, "el id es invalido")
				} else {
					if err := c.ShouldBindJSON(&req); err != nil {
						c.JSON(404, gin.H{
							"error": err.Error(),
						})
						return
					} else {

						if req.CodigoTransaccion == "" {
							c.JSON(401, gin.H{
								"error": "no se puede poner el codigo vacio",
							})

						} else {
							t, err := t.service.UpdateCod(int(id), req.CodigoTransaccion)
							if err != nil {
								c.JSON(404, gin.H{
									"error": err.Error(),
								})
								return
							}
							c.JSON(200, t)

						}

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

func (t *Transaccion) UpdateMon() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		token := c.GetHeader("token")

		if token != "" {
			if token == "39470939" {

				id, err := strconv.ParseInt(c.Param("id"), 10, 64)

				if err != nil {
					c.JSON(400, "el id es invalido")
				} else {
					if err := c.ShouldBindJSON(&req); err != nil {
						c.JSON(404, gin.H{
							"error": err.Error(),
						})
						return
					} else {

						if req.Monto == 0.0 {
							c.JSON(401, gin.H{
								"error": "no se puede poner el monto vacio",
							})

						} else {
							t, err := t.service.UpdateMon(int(id), req.Monto)
							if err != nil {
								c.JSON(404, gin.H{
									"error": err.Error(),
								})
								return
							}
							c.JSON(200, t)

						}

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
