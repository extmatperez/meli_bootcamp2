//EJERCICIO 1 AL 3 C1-GO WEB - TM

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

var productosSlice []Product

func BuscarPorID(ctxt *gin.Context) {
	se := false
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var productoEncontrado Product
		for _, e := range product {
			if strconv.Itoa(e.Id) == ctxt.Param("id") {
				productoEncontrado = e
				se = true
				break
			}
		}
		if se {
			ctxt.JSON(200, productoEncontrado)
		} else {
			ctxt.String(404, "No se encontro el empleado %s", ctxt.Param("id"))
		}
	}
}

func BuscarPorColor(ctxt *gin.Context) {
	parametro := ctxt.Query("Color")
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var filtrados []*Product
		for i, v := range product {
			if parametro == v.Color {
				filtrados = append(filtrados, &product[i])
			}
		}
		if len(filtrados) == 0 {
			ctxt.String(400, "No se encontró nada")
		} else {
			ctxt.JSON(200, filtrados)
		}
	}
}

func BuscarPorPrecio(ctxt *gin.Context) {
	parametro := ctxt.Query("Precio")
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var filtrados []*Product
		for i, v := range product {
			if parametro == v.Precio {
				filtrados = append(filtrados, &product[i])
			}
		}
		if len(filtrados) == 0 {
			ctxt.String(400, "No se encontró nada")
		} else {
			ctxt.JSON(200, filtrados)
		}
	}
}

func BuscarPorPublicado(ctxt *gin.Context) {
	parametro := ctxt.Param("publicado")
	se := false
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var productoEncontrado Product
		for _, e := range product {
			if strconv.FormatBool(e.Publicado) == parametro {
				productoEncontrado = e
				se = true
				break
			}
		}
		if se {
			ctxt.JSON(200, productoEncontrado)
		} else {
			ctxt.String(404, "No se encontro por publicado %s", ctxt.Param("Id"))
		}
	}
}

func mostrarTodo(ctx *gin.Context) {

}

func agregarEntidad(c *gin.Context) {
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var productos []Product //SLICE DEL JSON
		json.Unmarshal(data, &productos)
		var producto Product
		err := c.ShouldBindJSON(&producto)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(productos) == 0 {
			producto.Id = 1
		} else {
			producto.Id = productos[len(productos)-1].Id + 1
		}
		productos = append(productos, producto)

		actualizarJSON, err := json.Marshal(productos)
		err = os.WriteFile("./products.json", actualizarJSON, 0644)
		if err != nil {
			fmt.Println("No se pudo escribir")
		} else {
			fmt.Println("Grabo")
		}
		c.JSON(200, productos)
	}
}

func validarCampos(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != "2525" {
		ctx.JSON(401, gin.H{
			"error": "Token invalido, pruebe otro",
		})
	}
	var productoAux Product
	err := ctx.ShouldBindJSON(&productoAux)
	if err != nil {
		fmt.Println("error")
	} else {
		productoAux.Id = 1
		r := reflect.ValueOf(productoAux)
		for i := 0; i < r.NumField(); i++ {
			varValor := r.Field(i).Interface()
			s := reflect.TypeOf(varValor).Kind()
			if fmt.Sprint(s) == "string" {
				if varValor == "" {
					ctx.String(400, fmt.Sprintf("El campo %v no puede estar vacio", r.Type().Field(i).Name))
					return
				}
			} else {
				if varValor == 0 {
					ctx.String(400, fmt.Sprintf("El campo %v no puede ser cero", r.Type().Field(i).Name))
					return
				}
			}
		}

		ctx.JSON(200, gin.H{
			"todo OK": "Token y validacion correctas",
		})
	}
}

func main() {

	server := gin.Default()
	server.GET("/buscarId/:id", BuscarPorID)
	server.GET("/buscarPorColor", BuscarPorColor)
	server.GET("/buscarPorPrecio", BuscarPorPrecio)
	server.GET("/buscarPorPublicado/:publicado", BuscarPorPublicado)
	server.GET("/mostrarTodo")
	server.POST("/agregarProducto", agregarEntidad) //Agrega un producto de tipo Struct al .json
	server.POST("/validarCampos", validarCampos)
	server.Run()

}
