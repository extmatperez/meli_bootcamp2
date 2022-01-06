/* Ejercicio 1 - Estructura un JSON
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string),
fecha de transacción.
1) Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
2) Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.
*/

package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
}

/* func readData() []Users {

	var list []Users
	read_users, _ := os.ReadFile("./Users.json")

	if err := json.Unmarshal([]byte(read_users), &list); err != nil {
		log.Fatal(err)
	}
	return list
} */

/* Ejercicio 2 - Hola {nombre}

1) Crea dentro de la carpeta go-web un archivo llamado main.go
2) Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
3) Pegale al endpoint para corroborar que la respuesta sea la correcta.
*/

// Return Hello World!
func hello_world(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func hello_you(c *gin.Context) {
	name := c.Param("name") // c.GetQuery
	c.String(http.StatusOK, "Hello %s!", name)
}

// Return all users
func GetAll(c *gin.Context) {
	var users_list []Users
	read_users, err := os.ReadFile("./users.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Users not found!",
		})
		//panic(err)
	} else {
		json.Unmarshal(read_users, &users_list)
		c.JSON(http.StatusOK, gin.H{
			"users": users_list,
		})
	}
}

/* func getbyName(c *gin.Context) {

	var user_list = GetAll()
	var filtered []Users

	for _, us := range user_list {
		if c.Query("name") == us.Name {
			filtered = append(filtered, us)
		}
	}

	c.JSON(200, gin.H{
		"response": filtered,
	})

} */

/* Ejercicio 3 - Listar Entidad

Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
1) Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
2) Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
3) Genera un handler para el endpoint llamado “GetAll”.
4) Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
*/
func main() {
	router := gin.Default()
	// Return Hello World
	router.GET("/", hello_world)

	//Return name in params
	router.GET("/users/:name", hello_you)

	// Return all users
	router.GET("/users", GetAll)

	/* usersfiltered := router.Group("/usersfiltered")
	{
		usersfiltered.GET("/name", getbyName)
	} */

	router.Run()
}
