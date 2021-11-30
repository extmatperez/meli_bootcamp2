/*
Exercise 1:
Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
Luego genera la lógica de filtrado de nuestro array.
Devolver por el endpoint el array filtrado.
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
	// name := c.DefaultQuery("name", "Jose")
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Hello " + name,
	})
}

func saluteUser(c *gin.Context) {
	pUsers := GenerateUserList()
	num := c.Param("numUser")
	numFormat, _ := strconv.Atoi(num)
	fmt.Printf("num: %v - numFormat: %v\n", num, numFormat)
	name := string(pUsers[numFormat].FirstName)
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "hello " + name,
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

// func findQuery(ctx *gin.Context) {
// 	users := GenerateUserList()
// 	var filtrados []*User

// 	for i, v := range users {
// 		if ctx.Query("filtro") == v.Active {
// 			filtrados = append(filtrados, &users[i])
// 		}
// 	}

// 	if len(filtrados) == 0 {
// 		ctx.String(400, "No se encontró nada")
// 	} else {
// 		ctx.JSON(200, filtrados)
// 	}
// }

func main() {
	router := gin.Default()
	router.GET("/hello/:name", salute)
	router.GET("/helloUser/:numUser", saluteUser)
	router.GET("/users", getAllUsers)
	// router.GET("/findquery", findQuery)
	router.Run()
}
