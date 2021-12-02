/*Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.*/

/*Ejercicio 3
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
Genera un handler para el endpoint llamado “GetAll”.
Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Hola Fernando",
		})
	})
	router.GET("/usuarios", GetAll) /*  {
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Hola Fernando",
	}) */

	router.Run()
}

type Usuario struct {
	ID           int     `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creation_date"`
}

/* func leerUsuarios() []Usuario {
	jsonUsuarios := "/usuarios.json"
	data, err := os.ReadFile(jsonUsuarios)

	if err != nil {
		fmt.Printf("Error en la lectura %v", err)
	}

	return unmarshaleado(data)
}

func unmarshaleado(data []byte) []Usuario {
	var usuarios []Usuario

	if err := json.Unmarshal(data, &usuarios); err != nil {
		panic(err)
	}

	return usuarios
} */

func GetAll(c *gin.Context) {
	jsonUsuarios := "/Users/fgianni/Desktop/Fernando/repoMELI/meli_bootcamp2/6_goweb1/TM/go-web/usuarios.json"
	var usuarios []Usuario
	data, err := os.ReadFile(jsonUsuarios)

	if err != nil {
		fmt.Printf("Error en la lectura %v", err)
	}
	json.Unmarshal(data, &usuarios)

	c.JSON(200, usuarios)

	/* 	usuarios := leerUsuarios()
	   	c.JSON(200, gin.H{
	   		"Usuario": usuarios})
	*/
}
