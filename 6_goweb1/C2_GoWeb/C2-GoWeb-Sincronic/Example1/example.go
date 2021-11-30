package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

var users []User

func AddUser(ctx *gin.Context) {
	var us User

	errShoul := ctx.ShouldBind(&us)
	if errShoul != nil {
		ctx.JSON(400, gin.H{
			"Error": errShoul.Error(),
		})
		// ctx.String(400, "Se produjo un error: %v", errShoul.Error())
	} else {
		// us.ID = len(users) + 1
		if len(users) == 0 {
			us.ID = 1
		} else {
			us.ID = users[len(users)-1].ID + 1
		}
		users = append(users, us)
		ctx.JSON(200, us)
	}
}

func GetUsers(ctx *gin.Context) {

	token := ctx.GetHeader("token")

	if token != "" {
		if token == "123456" {
			if len(users) > 0 {
				ctx.JSON(200, users)
			} else {
				ctx.String(200, "No hay usuarios cargados en memoria")
			}
		} else {
			ctx.String(401, "Token incorrecto")
		}
	} else {
		ctx.String(401, "No ingreso un token")
	}

}

func main() {
	router := gin.Default()
	routerUsers := router.Group("/users")

	routerUsers.GET("", GetUsers)
	routerUsers.POST("/add", AddUser)
	router.Run()
}
