// Agregamos el package handler
package handler

// importo el package handler
import (
	"net/http"

	users "github.com/extmatperez/melibootcamp2/tree/montenegro_edgar/7_goweb2/afternoon_activities/Exercise_1_2/internal/users"
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
