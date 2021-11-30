package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Empleado struct {
	Nombre string `form:"name" json:"name"`
	Id     string `form:"id" json:"id"`
	Activo string `form:"active" json:"activa" binding:"required"`
}

// func BuscarEmpleado(ctxt *gin.Context) {
// 	parametro := ctx.Param("id")
// 	var empleado Empleado
// 	se := false
// 	for _,v := range parametro{
// 		if v.Id == parametro {
// 			emp = v
// 			se = true
// 			break
// 		}
// 	}
// 	if se {
// 		ctx.JSON(200,emp)
// 	}
// }

func Ejemplo(ctx *gin.Context) {
	contenido := ctx.Request.Body
	header := ctx.Request.Header
	metodo := ctx.Request.Method

	fmt.Println("Recibi algo: ")
	fmt.Println("Metodo: ", metodo)
	fmt.Println("Cabecera: ")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}

	fmt.Println("Contenido: ", contenido)
}

func main() {
	router := gin.Default()
	router.GET("/", Ejemplo)

	router.Run()
}
