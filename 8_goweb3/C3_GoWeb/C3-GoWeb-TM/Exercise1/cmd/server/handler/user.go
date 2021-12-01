package handler

import (
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-TM/Exercise1/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Height      int    `json:"height"`
	Active      bool   `json:"active"`
	CrationDate string `json:"cration_date"`
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
		// errLoad := us.service.LoadUser()
		// if errLoad != nil {
		// 	fmt.Printf("Error loading user")
		// } else {
		users, err := us.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error: %v", err)
		} else {
			ctx.JSON(200, users)
		}
		// }
	}
}

func (controller *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user request

		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona: %v", err)
		} else {
			response, err := controller.service.Store(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

func (controller *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "Error invalid id: %v", id)
			return
		}

		err = ctx.ShouldBindJSON(&req)

		if req.FirstName == "" {
			ctx.JSON(400, gin.H{"error": "First Name is required"})
			return
		}
		if req.LastName == "" {
			ctx.JSON(400, gin.H{"error": "Last Name is required"})
			return
		}
		if req.Email == "" {
			ctx.JSON(400, gin.H{"error": "Email is required"})
			return
		}
		if req.Age == 0 {
			ctx.JSON(400, gin.H{"error": "The Age cannot be zero"})
			return
		}
		if req.Height == 0 {
			ctx.JSON(400, gin.H{"error": "The Height cannot be zero"})
			return
		}
		if req.CrationDate == "" {
			ctx.JSON(400, gin.H{"error": "Creation date is required"})
			return
		}
		p, err := controller.service.Update(int(id), req.FirstName, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CrationDate)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)

		if err != nil {
			ctx.String(404, "Error in the body")
		} else {
			usuarioUpdate, err := controller.service.Update(int(id), req.FirstName, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CrationDate)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, usuarioUpdate)
			}
		}

	}
}
