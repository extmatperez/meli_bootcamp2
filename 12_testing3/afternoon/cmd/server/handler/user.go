package handler

import (
	"fmt"
	"os"
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/12_testing3/afternoon/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/12_testing3/afternoon/pkg/web"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Email    string  `json:"email"`
	Age      int     `json:"age"`
	Height   float64 `json:"height"`
	Active   bool    `json:"active"`
	Created  string  `json:"created"`
}

type User struct {
	service users.Service
}

func NewUser(ser users.Service) *User {
	return &User{service: ser}
}

// List users godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		users, err := u.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was an error: %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, users, ""))
		}
	}
}

// Create user godoc
// @Summary Create user
// @Tags Users
// @Description create user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param data body request true "The user data"
// @Success 200 {object} web.Response
// @Router /users [post]
func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		var user request
		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was an error when storing the user: %v", err)))
			return
		}

		newUser, errStore := u.service.Store(user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.Created)
		if errStore != nil {
			ctx.JSON(500, web.NewResponse(500, nil, fmt.Sprintf("There was an error when storing the user: %v", err)))
			return
		}

		ctx.JSON(200, web.NewResponse(200, newUser, ""))
	}
}

// Create user godoc
// @Summary Create user
// @Tags Users
// @Description create user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param data body request true "The user data"
// @Param id path int true "The user id"
// @Success 200 {object} web.Response
// @Router /users/{id} [put]
func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, errId := strconv.Atoi(ctx.Param("id"))
		if errId != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID"))
		}

		var user request
		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was an error when updating the user: %v", err)))
			return
		}

		checkMsg := validateUpdateFields(user)
		if checkMsg != "" {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Required field/s missing: %s", checkMsg)))
			return
		}

		updatedUser, errStore := u.service.Update(id, user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.Created)
		if errStore != nil {
			ctx.JSON(500, web.NewResponse(500, nil, fmt.Sprintf("There was an error when updating the user: %v", errStore)))
			return
		}

		if updatedUser == (users.User{}) {
			ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("User %d not found", id)))
		}

		ctx.JSON(200, web.NewResponse(200, updatedUser, ""))
	}
}

func (u *User) UpdateLastNameAge() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, errId := strconv.Atoi(ctx.Param("id"))
		if errId != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID"))
		}

		var user request
		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was an error when updating the user: %v", err)))
			return
		}

		checkMsg := validatePatchFields(user)
		if checkMsg != "" {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Required field/s missing: %s", checkMsg)))
			return
		}

		updatedUser, errStore := u.service.UpdateLastNameAge(id, user.LastName, user.Age)
		if errStore != nil {
			ctx.JSON(500, web.NewResponse(500, nil, fmt.Sprintf("There was an error when updating the user:  %v", err)))
			return
		}

		if updatedUser == (users.User{}) {
			ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("User %d not found", id)))
		}

		ctx.JSON(200, updatedUser)
	}
}

func (u *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid ID"))
			return
		}

		couldDelete, errDelete := u.service.Delete(id)
		if errDelete != nil {
			ctx.JSON(500, web.NewResponse(500, nil, fmt.Sprintf("There was an error when deleting the user:  %v", errDelete)))
			return
		}

		if !couldDelete {
			ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprintf("User %d not found", id)))
			return
		}

		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("The user %d has been deleted", id), ""))
	}
}

func validateUpdateFields(u request) string {
	msg := validatePatchFields(u)

	if u.Name == "" {
		msg = fmt.Sprintf("%s %s", msg, "name")
	}

	if u.Email == "" {
		msg = fmt.Sprintf("%s %s", msg, "email")
	}

	if u.Height == 0.0 {
		msg = fmt.Sprintf("%s %s", msg, "height")
	}

	if u.Created == "" {
		msg = fmt.Sprintf("%s %s", msg, "created")
	}

	return msg
}

func validatePatchFields(u request) string {
	msg := ""
	if u.LastName == "" {
		msg = fmt.Sprintf("%s %s", msg, "last_name")
	}

	if u.Age == 0 {
		msg = fmt.Sprintf("%s %s", msg, "age")
	}

	return msg
}

func validateToken(ctx *gin.Context) bool {
	var valid bool = true
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Invalid token"))
		valid = false
	}

	return valid
}
