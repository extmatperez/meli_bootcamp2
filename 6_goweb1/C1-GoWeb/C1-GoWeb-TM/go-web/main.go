/*
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
-> Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.
Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.

*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Height      int    `json:"height"`
	Active      bool   `json:"active"`
	CrationDate string `json:"cration_date"`
}

func (u *User) getFirstName() string {
	return u.FirstName
}

func salute(c *gin.Context) {
	name := c.Param("name")
	// name := c.DefaultQuery("name", "Jose")
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Hello " + name,
	})
}

func saluteUser(c *gin.Context) {
	data, err := os.ReadFile("users.json")
	var pUsersRead []User
	json.Unmarshal(data, &pUsersRead)
	if err != nil {
		fmt.Print("Error Read file")
	}
	num := c.Param("numUser")
	numFormat, _ := strconv.Atoi(num)
	fmt.Printf("num: %v - numFormat: %v\n", num, numFormat)
	name := string(pUsersRead[numFormat].FirstName)
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "hello " + name,
	})
}

func getAllUsers(c *gin.Context) {
	bytes, err := os.ReadFile("users.json")
	var pUsers []User
	json.Unmarshal(bytes, &pUsers)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error parseando el JSON de productos",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": pUsers,
	})
}

func main() {
	router := gin.Default()
	router.GET("/hello/:name", salute)
	router.GET("/users", getAllUsers)
	router.GET("/helloUser/:numUser", saluteUser)
	router.Run()
}
