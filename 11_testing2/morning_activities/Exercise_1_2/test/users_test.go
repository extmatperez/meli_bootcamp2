package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/11_testing2/morning_activities/Exercise_1_2/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/11_testing2/morning_activities/Exercise_1_2/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/11_testing2/morning_activities/Exercise_1_2/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/11_testing2/morning_activities/Exercise_1_2/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Agregamos el middleware del token
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token is required")
			return
		}
		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}

// Creamos nuestro server
func create_server() *gin.Engine {

	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")

	// Creamos una variable que guarde lo que vamos a escribir en nuestro json y se lo pasamos el repo (son las mismas de nuestro main)
	db := store.New(store.File_type, "./users_test.json") // Creo el json a la misma altura (copia de mi json)
	repo := users.New_repository(db)
	service := users.New_service(repo)
	controller := handler.New_user(service)

	router.Use(TokenAuthMiddleware())

	router.GET("/users", controller.Get_users())
	router.POST("/users", controller.Post_users())
	router.PUT("/users/:id", controller.Update_users())
	router.PATCH("/users/:id", controller.Update_users_first_name())
	router.DELETE("/users/:id", controller.Delete_users())

	// El router.Run() devuelve la función, pero nosotros solo necesitamos devolver las rutas
	return router
}

// Creamos la función que nos va a proporcionar request y response
func create_request_test(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

// Obtenemos los usuarios y validamos la respuesta
func Test_Get_users(t *testing.T) {
	router := create_server()
	req, rr := create_request_test(http.MethodGet, "/users", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var res web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Nil(t, err)
}
