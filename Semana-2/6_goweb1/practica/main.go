// Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
// 1. Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
// 2. Luego genera la lógica de filtrado de nuestro array.
// 3. Devolver por el endpoint el array filtrado.

// PRACTICA CON ARI

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
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Stock        int    `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creationDate"`
}

var products []Product

func getFiles(ctx *gin.Context) {

	data, err := os.ReadFile("../ejm/products.json") //traemos los archivos de products.json

	if err != nil {
		ctx.JSON(400, "ERROR AL IMPORTAR LOS DATOS")
		return // return para cortar la ejecucion
	}

	err = json.Unmarshal(data, &products)

	if err != nil {
		ctx.JSON(400, "ERROR AL UNMARSHALEAR") // partimos del error, si pasa es porque funciono
		return
	}

	ctx.JSON(200, products) //trajimos los products
}

func FiltrarNombre(ctx *gin.Context) {
	//ENVIO POR BODY
	var porBody Product
	err := ctx.ShouldBindJSON(&porBody)
	// con shouldBindJSON mandamos el JSON por el body y lo compara con la struct. EL BODY VIAJA DEL FRONT AL BACK
	if err != nil {
		ctx.JSON(400, "ERROR")
		return
	}

	fmt.Println(porBody)
	//ENVIO POR PARAM
	//todo lo que paso por param va a pasar por consola como STRING
	miFernet := ctx.Param("fernet")
	// el valor que pasa el front en el endpoint va a ser el valor de la clave fernet.
	fmt.Println(miFernet)
	miCoca, _ := strconv.Atoi(ctx.Param("coca"))

	fmt.Println(miCoca)
	fmt.Println(reflect.TypeOf(miCoca))

	//ENVIO POR QUERY-- es un modo de envio de informacion que se realiza a traves del endpoint luego del signo ?
	// por medio de una clave y un valor que nada tienen que ver con las etiquetas json.

	miProvincia := ctx.Query("provincia") // el nombre que le paso al query es el que puse en el front y tienen que ir en string
	miCiudad := ctx.Query("ciudad")

	fmt.Println(miProvincia, miCiudad)

}

func main() {
	router := gin.Default()
	router.GET("/products", getFiles)

	filtrar := router.Group("/filtrar")
	filtrar.GET("/nombre/", FiltrarNombre) // para query no ponemos nada, solo pasar clave y valor en la url

	filtrar.GET("/nombre/:fernet/:coca", FiltrarNombre)

	// al [pegarle al endpoint voy a reemplazar fernet y coca por las dos variables que quiera asignarle como valores

	router.Run()

}
