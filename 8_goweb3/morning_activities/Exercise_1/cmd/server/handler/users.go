// Agregamos el package handler
package handler

// importo el package handler
import (
	"net/http"
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/8_goweb3/morning_activities/Exercise_1/internal/users"
	"github.com/gin-gonic/gin"
)

// Creamos la request struct
type request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
}

// Creamos Users struct
type Users struct {
	service users.Service
}

// Agregar New_user function
func New_user(ser users.Service) *Users {
	return &Users{service: ser}
}

// Agregar Get_users handler que va a ser usado en el endpoint por main
func (us *Users) Get_users() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := us.service.Get_users()

		if err != nil {
			c.String(http.StatusBadRequest, "Something went wrong %v", err)
		} else {
			c.JSON(http.StatusOK, users)
		}
	}
}

// Agregar Post_users handler que va a ser usado en el endpoint por main
func (controller *Users) Post_users() gin.HandlerFunc {
	return func(c *gin.Context) {
		var use request

		err := c.ShouldBindJSON(&use)

		if err != nil {
			c.String(http.StatusBadRequest, "Something went wrong to post a new user %v", err)
		} else {
			response, err := controller.service.Post_users(use.FirstName, use.LastName, use.Email, use.Age, use.Height, use.Active, use.Date)
			if err != nil {
				c.String(http.StatusBadRequest, "Something went wrong to post a new user")
			} else {
				c.JSON(http.StatusOK, response)
			}
		}
	}
}

// Agregar Update_users handler que va a ser usado en el endpoint por main
func (controller *Users) Update_users() gin.HandlerFunc {
	return func(c *gin.Context) {
		var use request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.String(http.StatusBadRequest, "The id is not a valid id.")
		}

		err = c.ShouldBindJSON(&use)

		if err != nil {
			c.String(http.StatusBadRequest, "Something went wrong in body.")
		} else {
			user_updated, err := controller.service.Update_users(int(id), use.FirstName, use.LastName, use.Email, use.Age, use.Height, use.Active, use.Date)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			} else {
				c.JSON(http.StatusOK, user_updated)
			}
		}
	}
}

// Agregar Update_users_fields handler que va a ser usado en el endpoint por main
func (controller *Users) Update_users_first_name() gin.HandlerFunc {
	return func(c *gin.Context) {
		var use request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.String(http.StatusBadRequest, "The id is not a valid id.")
		}

		err = c.ShouldBindJSON(&use)

		if err != nil {
			c.String(http.StatusBadRequest, "Something went wrong in body.")
		} else {
			if use.FirstName == "" {
				c.String(http.StatusBadRequest, "The Name is required.")
				return
			}
			user_updated, err := controller.service.Update_users_first_name(int(id), use.FirstName)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			} else {
				c.JSON(http.StatusOK, user_updated)
			}
		}
	}
}
