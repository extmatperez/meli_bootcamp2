package main

import (
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
	token := validToken(ctx)
	if token {

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})

		} else {

			if len(users) == 0 {
				usr.ID = 1
			} else {
				usr.ID = users[len(users)-1].ID + 1
			}
			users = append(users, usr)
			ctx.JSON(200, usr)
		}
	}

}

func validToken(ctx *gin.Context) bool {

	token := ctx.GetHeader("token")

	if token != "" {
		if token == "123456" {
			return true
		} else {
			ctx.String(401, "no tiene permisos para realizar la petición solicitada")
			return false
		}
	} else {
		ctx.String(400, "No ingreso un token")
		return false
	}
}
