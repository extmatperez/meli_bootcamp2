package main

import (
	"encoding/json"
	"os"

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

// func filter(sliceUser []User, fields string, value string) []User {
// 	var filtrado []User

// 	var us User
// 	types := reflect.TypeOf(user)

// 	i := 0

// 	for i = 0; i < types.NumField(); i++ {
// 		if strings.ToLower(types.Field(i).Name) == fields {
// 			break
// 		}
// 	}
// 	for _, users := range sliceUser {
// 		if strings.Contains(reflect.ValueOf(users).Field(i).Interface) == value {
// 			filtrado = append(filtrado, users)
// 		}
// 	}

// 	return filtrado
// }

// func FilterUsers(ctx *gin.Context) {
// 	var tags []string
// 	tags = append(tags, "firs_name", "last_name", "age")
// 	// var searchedFields
// 	var usersFileters []User

// 	for _, tag := range tags {
// 		usersFileters = filter(users, tag)
// 	}

// 	if len(usersFileters) == 0 {
// 		ctx.String(200, "no hay coincidencias")
// 	} else {
// 		ctx.JSON(200, usersFileters)
// 	}
// }

func LoadData(ctx *gin.Context) {
	data, errReadJson := os.ReadFile("./users.json")
	if errReadJson != nil {
		ctx.String(400, "No se encontro el archivo")
	} else {
		json.Unmarshal(data, &users)
		ctx.String(200, "Usuarios cargados")
	}
}

func main() {
	router := gin.Default()
	routerUsers := router.Group("/users")

	routerUsers.GET("", GetUsers)
	routerUsers.POST("/add", AddUser)
	router.Run()
}
