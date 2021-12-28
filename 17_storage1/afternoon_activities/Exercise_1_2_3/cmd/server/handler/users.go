// Agregamos el package handler
package handler

// importo el package handler
import (
	"fmt"
	"net/http"

	/* "os" */
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/17_storage1/afternoon_activities/Exercise_1_2_3/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/17_storage1/afternoon_activities/Exercise_1_2_3/pkg/web"
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

// Agregamos una función para validar el token
/* func validate_token(c *gin.Context) bool {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, web.New_response(400, nil, "Missing token."))
		// c.String(http.StatusBadRequest, "Missing token.")
		return false
	}
	token_env := os.Getenv("TOKEN")
	if token != token_env {
		c.JSON(http.StatusBadRequest, web.New_response(404, nil, "Invalid Token."))
		// c.String(http.StatusBadRequest, "Invalid Token.")
		return false
	}
	return true
} */

// Store godoc
// @Summary Stores users
// @Tags Users
// @Description posts user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (us *Users) Get_users() gin.HandlerFunc {
	// Agregar Get_users handler que va a ser usado en el endpoint por main
	return func(c *gin.Context) {
		// Llamamos a la función validate_token
		/* if !validate_token(c) {
			return
		} */
		users, err := us.service.Get_users()

		if err != nil {
			c.JSON(http.StatusBadRequest, web.New_response(404, nil, fmt.Sprintf("Something went wrong %v", err)))
			// c.String(http.StatusBadRequest, "Something went wrong %v", err)
		} else {
			c.JSON(http.StatusOK, web.New_response(200, users, ""))
			// c.JSON(http.StatusOK, users)
		}
	}
}

// Post_users godoc
// @Summary Stores users
// @Tags Users
// @Description posts user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users/{id} [post]
func (controller *Users) Post_users() gin.HandlerFunc {
	// Agregar Post_users handler que va a ser usado en el endpoint por main
	return func(c *gin.Context) {
		// Llamamos a la función validate_token
		/* if !validate_token(c) {
			return
		} */
		var use request

		err := c.ShouldBindJSON(&use)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.New_response(404, nil, fmt.Sprintf("Something went wrong to post a new user %v", err)))
			// c.String(http.StatusBadRequest, "Something went wrong to post a new user %v", err)
		} else {
			response, err := controller.service.Post_users(use.FirstName, use.LastName, use.Email, use.Age, use.Height, use.Active, use.Date)
			if err != nil {
				c.JSON(http.StatusBadRequest, web.New_response(http.StatusBadRequest, nil, "Something went wrong to post a new user"))
				// c.String(http.StatusBadRequest, "Something went wrong to post a new user")
			} else {
				c.JSON(http.StatusOK, web.New_response(200, response, ""))
				// c.JSON(http.StatusOK, response)
			}
		}
	}
}

// Update_users godoc
// @Summary Edits a user
// @Tags Users
// @Description posts user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users/{id} [put]
func (controller *Users) Update_users() gin.HandlerFunc {
	// Agregar Update_users handler que va a ser usado en el endpoint por main
	return func(c *gin.Context) {
		// Llamamos a la función validate_token
		/* if !validate_token(c) {
			return
		} */
		var use request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.New_response(404, nil, "The id is not a valid id."))
			// c.String(http.StatusBadRequest, "The id is not a valid id.")
		}

		err = c.ShouldBindJSON(&use)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.New_response(http.StatusBadRequest, nil, "Something went wrong in body."))
			// c.String(http.StatusBadRequest, "Something went wrong in body.")
		} else {
			user_updated, err := controller.service.Update_users(int(id), use.FirstName, use.LastName, use.Email, use.Age, use.Height, use.Active, use.Date)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			} else {
				c.JSON(http.StatusOK, web.New_response(200, user_updated, ""))
				// c.JSON(http.StatusOK, user_updated)
			}
		}
	}
}

// Update_users_first_name godoc
// @Summary Edits users name
// @Tags Users
// @Description posts user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users/{id} [patch]
func (controller *Users) Update_users_first_name() gin.HandlerFunc {
	// Agregar Update_users_fields handler que va a ser usado en el endpoint por main
	return func(c *gin.Context) {
		// Llamamos a la función validate_token
		/* if !validate_token(c) {
			return
		} */
		var use request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.New_response(404, nil, "The id is not a valid id."))
			// c.String(http.StatusBadRequest, "The id is not a valid id.")
		}

		err = c.ShouldBindJSON(&use)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.New_response(http.StatusBadRequest, nil, "Something went wrong in body."))
			// c.String(http.StatusBadRequest, "Something went wrong in body.")
		} else {
			if use.FirstName == "" {
				c.JSON(http.StatusBadRequest, web.New_response(http.StatusBadRequest, nil, "The Name is required."))
				c.String(http.StatusBadRequest, "The Name is required.")
				return
			}
			user_updated, err := controller.service.Update_users_first_name(int(id), use.FirstName)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
			} else {
				c.JSON(http.StatusOK, web.New_response(200, user_updated, ""))
				// c.JSON(http.StatusOK, user_updated)
			}
		}
	}
}

// Delete godoc
// @Summary Deletes a user
// @Tags Users
// @Description Deletes user
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users/{id} [delete]
func (controller *Users) Delete_users() gin.HandlerFunc {
	// Agregar Delete_users handler que va a ser usado en el endpoint por main
	return func(c *gin.Context) {
		// Llamamos a la función validate_token
		/* if !validate_token(c) {
			return
		} */

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.New_response(404, nil, "The id is not a valid id."))
			// c.String(http.StatusBadRequest, "The id is not a valid id.")
		}

		err = controller.service.Delete_users(int(id))

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			c.JSON(http.StatusOK, web.New_response(201, nil, fmt.Sprintf("The user %v has been deleted.", id)))
			// c.String(http.StatusOK, "The user %v has been deleted.", id)
		}

	}
}
