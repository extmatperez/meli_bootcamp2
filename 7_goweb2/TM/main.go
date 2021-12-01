/*
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).
*/

/*
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo).
*/

/*
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::
1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”.
*/

package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

var productos []Producto

type Producto struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

func generarListaProductos() []Producto {

	data, _ := os.ReadFile("../productos.json")

	var lista []Producto

	json.Unmarshal(data, &lista)

	return lista
}

func insertarProducto(c *gin.Context) {

	token := c.GetHeader("token")

	validacionToken, err := validarToken(token)

	if !validacionToken {
		c.String(401, err)
	} else {
		var nuevoProducto Producto

		productos = generarListaProductos()

		err := c.ShouldBindJSON(&nuevoProducto)

		if err != nil {
			c.String(400, "Se produjo un error: %v", err.Error())
		} else {

			validacionCampos, err := validarCampos(nuevoProducto)

			if !validacionCampos {
				c.JSON(400, err)
			} else {
				if len(productos) == 0 {
					nuevoProducto.ID = 1
				} else {
					nuevoProducto.ID = productos[len(productos)-1].ID + 1
				}

				productos = append(productos, nuevoProducto)

				c.JSON(200, productos)
			}
		}
	}

}

func validarCampos(nuevoProducto Producto) (bool, string) {

	if nuevoProducto.Nombre == "" {
		return false, "El campo nombre está vacío"
	}
	if nuevoProducto.Color == "" {
		return false, "El campo color está vacío"
	}
	if nuevoProducto.Precio == "" {
		return false, "El campo precio está vacío"
	}
	if nuevoProducto.Stock == 0 {
		return false, "El campo stock está vacío"
	}
	if nuevoProducto.Codigo == "" {
		return false, "El campo código está vacío"
	}
	if nuevoProducto.FechaCreacion == "" {
		return false, "El campo fecha de creación está vacío"
	}

	return true, ""
}

func validarToken(token string) (bool, string) {

	if token != "" {
		if token == "789" {
			return true, ""
		} else {
			return false, "Token incorrecto"
		}
	} else {
		return false, "No se ingresó ningún token"
	}
}

func main() {

	router := gin.Default()

	router.POST("/productos/add", insertarProducto)

	router.Run("localhost:8080")

}
