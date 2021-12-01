package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        string `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

var users []User

func main() {

	router := gin.Default()
	router.POST("/add", AddUser)

	router.Run()

}

func AddUser(ctx *gin.Context) {
	var usr User
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})

	} else {
		if strings.Trim(usr.Nombre, " ") == "" {
			ctx.String(400, "error Nombre")
			return
		} else {
			if strings.Trim(usr.Apellido, " ") == "" {
				ctx.String(400, "error Apellido")
				return
			} else {
				if strings.Trim(usr.Activo, " ") == "" {
					ctx.String(400, "error Activo")
					return
				} else {

					if strings.Trim(usr.Email, " ") == "" {
						ctx.String(400, "error Email")
						return
					}
				}
			}
		}
	}

	if len(users) == 0 {
		usr.ID = 1
	} else {
		usr.ID = users[len(users)-1].ID + 1
	}
	users = append(users, usr)
	ctx.JSON(200, usr)
}
