package handler

import (
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-Sincronic/ExampleTM/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type User struct {
	service users.Service
}

func NewUser(ser users.Service) *User {
	return &User{
		service: ser}
}

func (us *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := us.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error: %v", err)
		} else {
			ctx.JSON(200, users)
		}
	}
}

func (controller *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user request

		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona: %v", err)
		} else {
			controller.service.Store(user.FirstName, user.LastName, user.Age)
		}
	}
}

func (controller *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "Error: Invalid ID")
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if req.FirstName == "" {
			ctx.JSON(400, gin.H{"error": "El nombre es requerido"})
			return
		}
		if req.LastName == "" {
			ctx.JSON(400, gin.H{"error": "El apellido es requerido"})
			return
		}
		if req.Age == 0 {
			ctx.JSON(400, gin.H{"error": "La edad es requerida"})
			return
		}
		us, err := controller.service.Update(int(id), req.FirstName, req.LastName, req.Age)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, us)
	}
}
