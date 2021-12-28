package handler

import (
	"bytes"
	"encoding/json"
	//"log"
	"net/http"
	"net/http/httptest"
	//"os"
	"testing"

	"github.com/gin-gonic/gin"
	personas "github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/personas"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/store"
	"github.com/stretchr/testify/assert"
)

/* func respondWithError(c *gin.Context, code int, message interface{}) {
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
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
} */

func createServer() *gin.Engine {
	r := gin.Default()

	db := store.New(store.FileType, "./personasTest.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := NewPersona(service)

	//r.Use(TokenAuthMiddleware())

	r.GET("/personas/get", controller.GetAll())
	r.POST("/personas/add", controller.Store())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllFuncional(t *testing.T) {
	router := createServer()

	req, res := createRequestTest(http.MethodGet, "/personas/get", "")

	router.ServeHTTP(res, req)
}

func TestStoreFuncional(t *testing.T) {
	router := createServer()

	var nuevaPersona request = request{Nombre: "JX", Apellido: "Rossi", Edad: 27}

	sliceDeBytes, _ := json.Marshal(nuevaPersona)

	req, res := createRequestTest(http.MethodPost, "/personas/add", string(sliceDeBytes))

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}