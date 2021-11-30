// WEB CONTEXT
// Request y Response. Lo que necesita el cliente y la respuesta del servidor respectivamente.

// Gin Context
// Nos permite pasar variables entre middleware, asi como headers, parametros, query string parameters, method entre otras variables del request.
// Se encarga de validar el json de una solicitud y generar una respuesta json.

// Se pueden agrupar endpoints que esten muy dentro de una URL, con server.Group("/raiz del grupo"), haciendo que desde dicha caracteristica sean accesibles.

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var empleados []Empleado = []Empleado{{"Rodrigo", "Activo", "1"}, {"Juampi", "Activo", "2"}, {"Florencia", "Inactivo", "3"}}

type Empleado struct {
	Nombre string `form:"name" json:"nombre"`
	Activo string `form:"activo" json:"activo"`
	Id     string `form:"id" json:"id" binding:"required"`
}

// Filtro de empleados.
func filtrar_empleados(ctx *gin.Context) {
	var filtrados []*Empleado

	for _, v := range empleados {
		// Aqui va el nombre de la columna de como se llama en la estructura.
		if ctx.Query("filtro") == v.Activo {
			filtrados = append(filtrados, &v)
		}
	}

	ctx.JSON(200, gin.H{
		"empleados": filtrados,
	})
}

func buscar_empleado(ctx *gin.Context) {
	parametro := ctx.Param(("id"))
	var emp Empleado
	se := false
	for _, v := range empleados {
		if v.Id == parametro {
			emp = v
			se = true
			break
		}
	}

	if se {
		ctx.JSON(200, emp)
	} else {
		ctx.String(404, "No se encuentra el empleado %s", parametro)
	}
}

func ejemplo(ctx *gin.Context) {
	contenido := ctx.Request.Body
	header := ctx.Request.Header
	metodo := ctx.Request.Method

	fmt.Println("Recibi algo!")
	fmt.Println("Metodo: ", metodo)
	fmt.Println("Cabecera: ")

	for k, v := range header {
		fmt.Println(k, " : ", v)
	}

	fmt.Println("Contenido", contenido)
	// El context te permite formatear la salida! Es un formato que se puede concatenar. Todas las veces que se llame se va a ir agregando el contenido.
	ctx.String(200, "Termine %.2f", 2.6666)
}

func main() {
	router := gin.Default()

	router.GET("/ejemplo", ejemplo)
	router.GET("/buscar_empleado/:id", buscar_empleado)
	router.GET("/filtrar_empleados", filtrar_empleados)

	router.Run()
}
