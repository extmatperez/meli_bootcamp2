/*
Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática.
Utilizando path parameters el endpoint debería ser /temática/:id (recuerda que siempre tiene
que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.
1- Genera una nueva ruta.
2- Genera un handler para la ruta creada.
3- Dentro del handler busca el item que necesitas.
4- Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID          int    `form:"id" json:"id"`
	FirstName   string `form:"firs_name" json:"first_name"`
	LastName    string `form:"last_name" json:"last_name"`
	Email       string `form:"email" json:"email"`
	Age         int    `form:"age" json:"age"`
	Height      int    `form:"height" json:"height"`
	Active      bool   `form:"active" json:"active"`
	CrationDate string `form:"create_date" json:"cration_date"`
}

func (u *User) getFirstName() string {
	return u.FirstName
}

func salute(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Hello " + name,
	})
}

func getUserById(c *gin.Context) {
	allUsers := GenerateUserList()
	var pUsersRead []User

	idParam := c.Param("id")
	lastId := 0

	for i := 0; i < len(allUsers); i++ {
		if idParam != "" {
			id, errAtoi := strconv.Atoi(idParam)
			if errAtoi == nil {
				if allUsers[i].ID == id {
					lastId = i
					fmt.Println("Encontrado ID", lastId)
				}
			}
		}
	}

	if lastId != 0 {
		fmt.Println("Last ID", lastId)
		pUsersRead = append(pUsersRead, allUsers[lastId])
	} else {
		c.String(404, "Usuario con el ID %v no encontrado", idParam)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": pUsersRead,
	})
}

func getAllUsers(ctx *gin.Context) {
	allUsers := GenerateUserList()

	var pUsersRead []User

	firstNameQuery := ctx.Query("first_name")
	lastNameQuery := ctx.Query("last_name")
	idQuery := ctx.Query("id")
	emailQuery := ctx.Query("email")
	ageQuery := ctx.Query("age")
	heightQuery := ctx.Query("height")
	activeQuery := ctx.Query("active")
	crationDateQuery := ctx.Query("creation_date")

	for i := 0; i < len(allUsers); i++ {
		includesAllFilters := true
		if idQuery != "" {
			id, err := strconv.Atoi(idQuery)
			if err == nil {
				if allUsers[i].ID != id {
					includesAllFilters = false
				}
			}
		}
		if firstNameQuery != "" {
			if !strings.EqualFold(allUsers[i].FirstName, firstNameQuery) {
				includesAllFilters = false
			}
		}
		if lastNameQuery != "" {
			if !strings.EqualFold(allUsers[i].LastName, lastNameQuery) {
				includesAllFilters = false
			}
		}
		if emailQuery != "" {
			if !strings.EqualFold(allUsers[i].Email, emailQuery) {
				includesAllFilters = false
			}
		}
		if ageQuery != "" {
			age, err := strconv.Atoi(ageQuery)
			if err == nil {
				if allUsers[i].Age != age {
					includesAllFilters = false
				}
			}
		}
		if heightQuery != "" {
			height, err := strconv.Atoi(heightQuery)
			if err == nil {
				if allUsers[i].Height != height {
					includesAllFilters = false
				}
			}
		}
		if activeQuery != "" {
			if activeQuery == "true" || activeQuery == "false" {
				var active bool
				if activeQuery == "true" {
					active = true
				} else if activeQuery == "false" {
					active = false
				}

				if allUsers[i].Active != active {
					includesAllFilters = false
				}
			}
		}
		if crationDateQuery != "" {
			if !strings.EqualFold(allUsers[i].CrationDate, crationDateQuery) {
				includesAllFilters = false
			}
		}
		if includesAllFilters {
			pUsersRead = append(pUsersRead, allUsers[i])
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": pUsersRead,
	})
}

func GenerateUserList() []User {
	data, err := os.ReadFile("users.json")
	var pUsersRead []User

	if err != nil {
		fmt.Println("Error al encontrar el archivo json")
		return pUsersRead
	}

	errUnmarshall := json.Unmarshal(data, &pUsersRead)
	if errUnmarshall != nil {
		fmt.Println("Error parseando el JSON de productos")
		return pUsersRead
	}
	return pUsersRead
}

func main() {
	router := gin.Default()
	router.GET("/hello/:name", salute)
	router.GET("/users/:id", getUserById)
	router.GET("/users", getAllUsers)
	// router.GET("/findquery", findQuery)
	router.Run()
}
