/* Ejercicio 1 - Crear Entidad
Se debe implementar la funcionalidad para crear la entidad. para eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global). */

/* Ejercicio 2 - Validación de campos
Se debe implementar las validaciones de los campos al momento de enviar la petición, para
eso se deben seguir los siguientes pasos:
1. Se debe validar todos los campos enviados en la petición, todos los campos son
requeridos
2. En caso que algún campo no esté completo se debe retornar un código de error 400
con el mensaje “el campo %s es requerido”.
(En %s debe ir el nombre del campo que no está completo). */

/* Ejercicio 3 - Validar Token
Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
deben seguir los siguientes pasos::

1. Al momento de enviar la petición se debe validar que un token sea enviado
2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
mensaje que “no tiene permisos para realizar la petición solicitada”. */

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

var user []Users

// Return all users
func get_users(c *gin.Context) {
	if len(user) == 0 {
		read_users, _ := os.ReadFile("./users.json")
		err := json.Unmarshal(read_users, &user)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Users not found!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"users": user,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": user,
		})
	}
}
func post_users(c *gin.Context) {
	var user_id Users

	err := c.ShouldBindJSON(&user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		last_id := user[len(user)-1].ID + 1
		user_id.ID = last_id
		user = append(user, user_id)
		c.JSON(http.StatusOK, user_id)
	}
}

func main() {
	router := gin.Default()
	router.GET("/users", get_users)
	router.POST("/users", post_users)

	router.Run()
}
