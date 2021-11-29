/*
Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.
*/

/*
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
Genera un handler para el endpoint llamado “GetAll”.
Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
*/

package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

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

func saludar(c *gin.Context) {

	// data, _ := os.ReadFile("./productos.json")

	// var lista []Producto

	// json.Unmarshal(data, &lista)

	// c.JSON(200, gin.H{
	// 	"message": "Hola " + lista[0].Nombre,
	// })

	c.JSON(200, gin.H{
		"message": "Hola Benjamin",
	})
}

func GetAll(c *gin.Context) {

	data, _ := os.ReadFile("./productos.json")

	var lista []Producto

	json.Unmarshal(data, &lista)

	c.JSON(200, gin.H{
		"productos": lista,
	})
}

func main() {

	router := gin.Default()

	router.GET("/hola", saludar)

	router.GET("/productos", GetAll)

	router.Run()
}
