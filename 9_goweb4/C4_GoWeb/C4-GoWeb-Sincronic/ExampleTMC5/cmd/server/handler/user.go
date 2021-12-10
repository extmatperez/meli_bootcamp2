package handler

import (
	"fmt"
	"os"
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTMC5/internal/users"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTMC5/pkg/web"
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

func ValidateToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")

	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "Token not found"))
		return false
	}

	tokenENV := os.Getenv("TOKEN")

	if token != tokenENV {
		ctx.JSON(400, web.NewResponse(400, nil, "Invalid token"))
		return false
	}

	return true
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (us *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ValidateToken(ctx) {
			return
		}

		users, err := us.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was a mistake: %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, users, ""))
		}
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (controller *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidateToken(ctx) {
			return
		}

		var user request

		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.String(400, "There was an error wanting to load a user: %v", err)
		} else {
			response, err := controller.service.Store(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate)
			if err != nil {
				ctx.String(400, "Could not load user %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

func (controller *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidateToken(ctx) {
			return
		}
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

func (controller *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidateToken(ctx) {
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("The user %d is deleted", id)})
	}
}

func (controller *User) UpdateLastName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidateToken(ctx) {
			return
		}

		var req request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(404, gin.H{"error in the body": err.Error()})
		} else {
			if req.LastName == "" {
				ctx.JSON(400, gin.H{"error": "The LastName is required"})
				return
			}
			userUpdate, err := controller.service.UpdateLastName(int(id), req.LastName)
			if err != nil {
				ctx.JSON(404, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, userUpdate)
		}
	}
}

func (controller *User) UpdateAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		err = ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(404, gin.H{"error in the body": err.Error()})
		} else {
			if req.Age == 0 {
				ctx.JSON(400, gin.H{"error": "The Age cannot be zero"})
				return
			}
			userUpdate, err := controller.service.UpdateAge(int(id), req.Age)
			if err != nil {
				ctx.JSON(404, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, userUpdate)
		}
	}
}
