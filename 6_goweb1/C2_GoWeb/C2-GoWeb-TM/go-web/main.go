/*
Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
siguientes pasos:
1. Crea un endpoint mediante POST el cual reciba la entidad.
2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
deberán ir guardando todas las peticiones que se vayan realizando.
3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
nuevo registro (sin tener una variable de último ID a nivel global).
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

var userList = []User{}

func GenerateUserList() []User {
	bytes, err := os.ReadFile("usezrs.json")
	var pUsersRead []User

	if err != nil {
		fmt.Println("Error al encontrar el archivo json")
		return pUsersRead
	}

	errUnmarshall := json.Unmarshal(bytes, &pUsersRead)
	if errUnmarshall != nil {
		fmt.Println("Error parseando el JSON de productos")
		return pUsersRead
	}
	return pUsersRead
}

func salute(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Hello " + name,
	})
}

func saluteUser(c *gin.Context) {
	num := c.Param("numUser")
	numFormat, _ := strconv.Atoi(num)
	name := string(userList[numFormat].FirstName)
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "hello " + name,
	})
}

func getAllUsers(ctx *gin.Context) {
	var pUsersRead []User

	firstNameQuery := ctx.Query("first_name")
	lastNameQuery := ctx.Query("last_name")
	idQuery := ctx.Query("id")
	emailQuery := ctx.Query("email")
	ageQuery := ctx.Query("age")
	heightQuery := ctx.Query("height")
	activeQuery := ctx.Query("active")
	crationDateQuery := ctx.Query("creation_date")

	if len(userList) > 0 {
		for i := 0; i < len(userList); i++ {
			includesAllFilters := true
			if idQuery != "" {
				id, err := strconv.Atoi(idQuery)
				if err == nil {
					if userList[i].ID != id {
						includesAllFilters = false
					}
				}
			}
			if firstNameQuery != "" {
				if !strings.EqualFold(userList[i].FirstName, firstNameQuery) {
					includesAllFilters = false
				}
			}
			if lastNameQuery != "" {
				if !strings.EqualFold(userList[i].LastName, lastNameQuery) {
					includesAllFilters = false
				}
			}
			if emailQuery != "" {
				if !strings.EqualFold(userList[i].Email, emailQuery) {
					includesAllFilters = false
				}
			}
			if ageQuery != "" {
				age, err := strconv.Atoi(ageQuery)
				if err == nil {
					if userList[i].Age != age {
						includesAllFilters = false
					}
				}
			}
			if heightQuery != "" {
				height, err := strconv.Atoi(heightQuery)
				if err == nil {
					if userList[i].Height != height {
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

					if userList[i].Active != active {
						includesAllFilters = false
					}
				}
			}
			if crationDateQuery != "" {
				if !strings.EqualFold(userList[i].CrationDate, crationDateQuery) {
					includesAllFilters = false
				}
			}
			if includesAllFilters {
				pUsersRead = append(pUsersRead, userList[i])
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"users": pUsersRead,
		})
	} else {
		ctx.String(200, "No hay usuarios cargados en memoria")
	}
}

func AddUser(ctx *gin.Context) {
	var us User

	errShoul := ctx.ShouldBind(&us)
	if errShoul != nil {
		ctx.JSON(400, gin.H{
			"Error": errShoul.Error(),
		})
	} else {
		if len(userList) == 0 {
			us.ID = 1
		} else {
			us.ID = userList[len(userList)-1].ID + 1
		}
		userList = append(userList, us)
		ctx.JSON(200, us)
	}
}

func main() {
	userList = GenerateUserList()
	router := gin.Default()

	routerSalute := router.Group("/hello")
	routerSalute.GET("/:name", salute)
	routerSalute.GET("/user/:numUser", saluteUser)

	routerUsers := router.Group("/users")
	routerUsers.GET("", getAllUsers)
	routerUsers.POST("/add", AddUser)

	router.Run()
}
