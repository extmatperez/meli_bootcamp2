package handler

import (
	"fmt"
	"os"
	"strconv"

	productos "github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

type Producto struct {
	service productos.Service
}

func NewProducto(s productos.Service) *Producto {
	return &Producto{service: s}
}

func (controller *Producto) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			//c.String(400, "Token no enviado")
			c.JSON(400, web.NewResponse(400, nil, "Token no enviado"))
			return
		}

		tokenENV := os.Getenv("TOKEN")

		if token != tokenENV {
			//c.String(401, "Token inválido")
			c.JSON(400, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		productos, err := controller.service.GetAll()

		if err != nil {
			//c.String(400, "Hubo un error: %v", err.Error())
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error: %v", err)))
		} else {
			//c.JSON(200, productos)
			c.JSON(200, web.NewResponse(200, productos, ""))
		}

	}
}

func (controller *Producto) Store() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			//c.String(400, "Token no enviado")
			c.JSON(400, web.NewResponse(400, nil, "Token no enviado"))
			return
		}

		tokenENV := os.Getenv("TOKEN")

		if token != tokenENV {
			//c.String(401, "Token inválido")
			c.JSON(400, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		var nuevoProducto request

		err := c.ShouldBindJSON(&nuevoProducto)

		if err != nil {
			//c.String(400, "Hubo un error al recibir los datos en el body: %v", err.Error())
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al recibir los datos en el body: %v", err.Error())))
		} else {

			if nuevoProducto.Nombre == "" {
				//c.String(400, "Error: el nombre es obligatorio")
				c.JSON(400, web.NewResponse(400, nil, "Error: el nombre es obligatorio"))
				return
			}

			if nuevoProducto.Color == "" {
				//c.String(400, "Error: el color es obligatorio")
				c.JSON(400, web.NewResponse(400, nil, "Error: el color es obligatorio"))
				return
			}

			if nuevoProducto.Precio == "" {
				//c.String(400, "Error: el precio es obligatorio")
				c.JSON(400, web.NewResponse(400, nil, "Error: el precio es obligatorio"))
				return
			}

			if nuevoProducto.Codigo == "" {
				//c.String(400, "Error: el stock es obligatorio")
				c.JSON(400, web.NewResponse(400, nil, "Error: el stock es obligatorio"))
				return
			}

			if nuevoProducto.FechaCreacion == "" {
				//c.String(400, "Error: la fecha de cración es obligatoria")
				c.JSON(400, web.NewResponse(400, nil, "Error: la fecha de cración es obligatoria"))
				return
			}

			response, err := controller.service.Store(nuevoProducto.Nombre, nuevoProducto.Color, nuevoProducto.Precio, nuevoProducto.Stock, nuevoProducto.Codigo, nuevoProducto.Publicado, nuevoProducto.FechaCreacion)

			if err != nil {
				//c.String(400, "Hubo un error al guardar los datos: %v", err.Error())
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al guardar los datos: %v", err.Error())))
			} else {
				//c.JSON(200, response)
				c.JSON(200, web.NewResponse(200, response, ""))
			}
		}

	}
}

func (controller *Producto) Update() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			//c.String(400, "Token no enviado")
			c.JSON(400, web.NewResponse(400, nil, "Token no enviado"))
			return
		}

		tokenENV := os.Getenv("TOKEN")

		if token != tokenENV {
			//c.String(401, "Token inválido")
			c.JSON(400, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		var productoActualizado request

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			//c.String(400, "No se pudo leer el id por parámetro")
			c.JSON(400, web.NewResponse(400, nil, "No se pudo leer el id por parámetro"))
		} else {

			err = c.ShouldBindJSON(&productoActualizado)

			if err != nil {
				//c.String(400, "Hubo un error al recibir los datos en el body: %v", err.Error())
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al recibir los datos en el body: %v", err.Error())))
			} else {

				if productoActualizado.Nombre == "" {
					//c.String(400, "Error: el nombre es obligatorio")
					c.JSON(400, web.NewResponse(400, nil, "Error: el nombre es obligatorio"))
					return
				}

				if productoActualizado.Color == "" {
					//c.String(400, "Error: el color es obligatorio")
					c.JSON(400, web.NewResponse(400, nil, "Error: el color es obligatorio"))
					return
				}

				if productoActualizado.Precio == "" {
					//c.String(400, "Error: el precio es obligatorio")
					c.JSON(400, web.NewResponse(400, nil, "Error: el precio es obligatorio"))
					return
				}

				if productoActualizado.Codigo == "" {
					//c.String(400, "Error: el stock es obligatorio")
					c.JSON(400, web.NewResponse(400, nil, "Error: el stock es obligatorio"))
					return
				}

				if productoActualizado.FechaCreacion == "" {
					//c.String(400, "Error: la fecha de cración es obligatoria")
					c.JSON(400, web.NewResponse(400, nil, "Error: la fecha de cración es obligatoria"))
					return
				}

				productoActualizado, err := controller.service.Update(id, productoActualizado.Nombre, productoActualizado.Color, productoActualizado.Precio, productoActualizado.Stock, productoActualizado.Codigo, productoActualizado.Publicado, productoActualizado.FechaCreacion)

				if err != nil {
					//c.String(404, err.Error())
					c.JSON(404, web.NewResponse(404, nil, err.Error()))
				} else {
					//c.JSON(200, productoActualizado)
					c.JSON(200, web.NewResponse(200, productoActualizado, ""))
				}

			}
		}
	}
}

func (controller *Producto) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			//c.String(400, "Token no enviado")
			c.JSON(400, web.NewResponse(400, nil, "Token no enviado"))
			return
		}

		tokenENV := os.Getenv("TOKEN")

		if token != tokenENV {
			//c.String(401, "Token inválido")
			c.JSON(400, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			//c.String(400, "No se pudo leer el id por parámetro")
			c.JSON(400, web.NewResponse(400, nil, "No se pudo leer el id por parámetro"))
		} else {

			err := controller.service.Delete(id)

			if err != nil {
				//c.String(404, err.Error())
				c.JSON(404, web.NewResponse(404, nil, err.Error()))
			} else {
				//c.String(200, "El producto %v fue eliminado", id)
				c.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %v fue eliminado", id), ""))
			}

		}
	}
}

func (controller *Producto) UpdateNombrePrecio() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			//c.String(400, "Token no enviado")
			c.JSON(400, web.NewResponse(400, nil, "Token no enviado"))
			return
		}

		tokenENV := os.Getenv("TOKEN")

		if token != tokenENV {
			//c.String(401, "Token inválido")
			c.JSON(400, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		var productoActualizado request

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			//c.String(400, "No se pudo leer el id por parámetro")
			c.JSON(400, web.NewResponse(400, nil, "No se pudo leer el id por parámetro"))
		} else {

			err = c.ShouldBindJSON(&productoActualizado)

			if err != nil {
				//c.String(400, "Hubo un error al recibir los datos en el body: %v", err.Error())
				c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al recibir los datos en el body: %v", err.Error())))
			} else {

				if productoActualizado.Nombre == "" {
					//c.String(400, "Error: el nombre es obligatorio")
					c.JSON(400, web.NewResponse(400, nil, "Error: el nombre es obligatorio"))
					return
				}

				if productoActualizado.Precio == "" {
					//c.String(400, "Error: el precio es obligatorio")
					c.JSON(400, web.NewResponse(400, nil, "Error: el precio es obligatorio"))
					return
				}

				productoActualizado, err := controller.service.UpdateNombrePrecio(id, productoActualizado.Nombre, productoActualizado.Precio)

				if err != nil {
					//c.String(404, err.Error())
					c.JSON(404, web.NewResponse(404, nil, err.Error()))
				} else {
					//c.JSON(200, productoActualizado)
					c.JSON(200, web.NewResponse(200, productoActualizado, ""))
				}

			}
		}
	}
}
