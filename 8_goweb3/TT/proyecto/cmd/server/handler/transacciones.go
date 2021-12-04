package handler

import (
	"fmt"
	"strconv"

	transacciones "github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/pkg/web"
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

		t, err := t.service.Load()

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			return
		}
		c.JSON(200, web.NewResponse(200, t, ""))

	}
}

func (t *Transaccion) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		t, err := t.service.GetAll()

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			return
		}
		c.JSON(200, web.NewResponse(200, t, ""))

	}
}

func (t *Transaccion) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request

		if err := c.Bind(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			return
		} else {
			switch {
			case req.Monto == 0.0:
				c.JSON(401, web.NewResponse(400, nil, "no se puede poner el monto vacio"))
			case req.Emisor == "":
				c.JSON(401, web.NewResponse(400, nil, "no se puede emitir al emisor"))
			case req.Moneda == "":
				c.JSON(401, web.NewResponse(400, nil, "no se puede omitir el tipo de moneda"))
			case req.Receptor == "":
				c.JSON(401, web.NewResponse(400, nil, "no se puede emitir al receptor"))
			default:
				t, err := t.service.Store(req.ID, req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaCreacion)
				if err != nil {
					c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
					return
				}
				c.JSON(200, web.NewResponse(200, t, ""))

			}

		}

	}
}

func (t *Transaccion) FindById() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, "el id es invalido"))
		} else {

			t, err := t.service.FindById(int(id))
			if err != nil {
				c.JSON(404, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(200, web.NewResponse(200, t, ""))

		}
	}
}

func (t *Transaccion) FilterBy() gin.HandlerFunc {
	return func(c *gin.Context) {

		moneda := c.Query("moneda")
		emisor := c.Query("emisor")
		receptor := c.Query("receptor")
		fechacreacion := c.Query("fechacreacion")
		codigotransaccion := c.Query("codigotransaccion")

		t, err := t.service.FilterBy(moneda, emisor, receptor, fechacreacion, codigotransaccion)
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			return
		}
		c.JSON(200, web.NewResponse(200, t, ""))

	}
}

func (t *Transaccion) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(400, "el id es invalido")
		} else {
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
				return
			} else {
				switch {
				case req.Monto == 0.0:
					c.JSON(401, web.NewResponse(400, nil, "no se puede poner el monto vacio"))
				case req.Emisor == "":
					c.JSON(401, web.NewResponse(400, nil, "no se puede emitir al emisor"))
				case req.Moneda == "":
					c.JSON(401, web.NewResponse(400, nil, "no se puede omitir el tipo de moneda"))
				case req.Receptor == "":
					c.JSON(401, web.NewResponse(400, nil, "no se puede emitir al receptor"))
				default:
					t, err := t.service.Update(int(id), req.CodigoTransaccion, req.Moneda, req.Monto, req.Emisor, req.Receptor, req.FechaCreacion)
					if err != nil {
						c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
						return
					}
					c.JSON(200, web.NewResponse(200, t, ""))

				}

			}

		}
	}
}

func (t *Transaccion) UpdateCod() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(401, web.NewResponse(400, nil, "el id es invalido"))
		} else {
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
				return
			} else {

				if req.CodigoTransaccion == "" {
					c.JSON(401, web.NewResponse(400, nil, "no se puede dejar el codigo vacio"))

				} else {
					t, err := t.service.UpdateCod(int(id), req.CodigoTransaccion)
					if err != nil {
						c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
						return
					}
					c.JSON(200, web.NewResponse(200, t, ""))

				}

			}

		}
	}
}

func (t *Transaccion) UpdateMon() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(401, web.NewResponse(400, nil, "el id es invalido"))
		} else {
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
				return
			} else {

				if req.Monto == 0.0 {
					c.JSON(401, web.NewResponse(400, nil, "no se puede poner el monto vacio"))

				} else {
					t, err := t.service.UpdateMon(int(id), req.Monto)
					if err != nil {
						c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
						return
					}
					c.JSON(200, web.NewResponse(200, t, ""))

				}

			}

		}
	}
}

func (t *Transaccion) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(401, web.NewResponse(400, nil, "el id es invalido"))
		} else {

			err := t.service.Delete(int(id))
			if err != nil {
				c.JSON(201, web.NewResponse(201, nil, fmt.Sprintf("%v", err)))
				return
			}
		}
	}
}
