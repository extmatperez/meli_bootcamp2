package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	productos "github.com/extmatperez/meli_bootcamp2/tree/parra_diego/12_testing3/TT/ejercicio_2/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/12_testing3/TT/ejercicio_2/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/12_testing3/TT/ejercicio_2/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
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
}

func createServer() *gin.Engine {
	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./personasSalidaTest.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	controller := NewProduct(service)

	router.Use(TokenAuthMiddleware())

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	//
	router.PUT("/products/:id", controller.Update())
	router.DELETE("/products/:id", controller.Delete())
	router.PATCH("/products/:id", controller.UpdateNamePrice())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_UpdatePersonas(t *testing.T) {
	router := createServer()

	nuevoProducto := `{	"id":1,
	"name":"neve",
	"color":"azul",
	"price":34,
	"stock":3,
	"code":"e3",
	"publish":true,
	"date":"33/7"
	}`

	req, rr := createRequestTest(http.MethodPut, "/products/1", nuevoProducto)

	router.ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Code)

}

func Test_Delete(t *testing.T) {
	router := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/products/:id", "")

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)

	assert.Equal(t, 200, respuesta.Code)
	assert.Nil(t, err)
}
