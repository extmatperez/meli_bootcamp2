package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/TM/go-web/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/11_testing2/TM/go-web/internal/products"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/TM/go-web/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/TM/go-web/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Falta Token"))
		} else if token != os.Getenv("TOKEN") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "No tiene permisos para realizar la petición solicitada"))
		}
		c.Next()
	}
}
func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "1234")
	db := store.New("file", "./productsTest.json")
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	controller := handler.NewProduct(service)

	router := gin.Default()
	router.Use(TokenAuthMiddleware())
	productsRoute := router.Group("products")
	productsRoute.GET("", controller.GetAll())
	productsRoute.GET("/filter", controller.Filter())
	productsRoute.GET("/:id", controller.FindById())
	productsRoute.POST("", controller.Store())
	productsRoute.PUT("/:id", controller.Update())
	productsRoute.DELETE("/:id", controller.Delete())
	productsRoute.PATCH("/:id", controller.UpdateNameAndPrice())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "1234")

	return req, httptest.NewRecorder()
}

func TestUpdate(t *testing.T) {
	router := createServer()
	updateRequest := handler.Request{
		Nombre:        "Wine Change Name",
		Color:         "Pink Change",
		Precio:        269,
		Stock:         38,
		Codigo:        "3C3CFFJH0ET601111",
		Publicado:     false,
		FechaCreacion: "26/08/2021",
	}

	dataByte, _ := json.Marshal(updateRequest)

	req, rr := createRequestTest(http.MethodPut, fmt.Sprintf("/products/%d", 1), string(dataByte))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	updateRequest = handler.Request{
		Nombre:        "Wine Change Name 1",
		Color:         "Pink Change 1",
		Precio:        270,
		Stock:         40,
		Codigo:        "3C3CFFJH0ET60xxxx",
		Publicado:     false,
		FechaCreacion: "26/08/2021",
	}

	dataByte, _ = json.Marshal(updateRequest)

	req, rr = createRequestTest(http.MethodPut, fmt.Sprintf("/products/%d", 1), string(dataByte))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateNotFound(t *testing.T) {
	router := createServer()
	updateRequest := handler.Request{
		Nombre:        "Wine Change Name",
		Color:         "Pink Change",
		Precio:        269,
		Stock:         38,
		Codigo:        "3C3CFFJH0ET601111",
		Publicado:     false,
		FechaCreacion: "26/08/2021",
	}

	dataByte, _ := json.Marshal(updateRequest)

	req, rr := createRequestTest(http.MethodPut, fmt.Sprintf("/products/%d", -1), string(dataByte))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 404, rr.Code)
}
func TestUpdateInvalidId(t *testing.T) {
	router := createServer()
	updateRequest := handler.Request{
		Nombre:        "Wine Change Name",
		Color:         "Pink Change",
		Precio:        269,
		Stock:         38,
		Codigo:        "3C3CFFJH0ET601111",
		Publicado:     false,
		FechaCreacion: "26/08/2021",
	}

	dataByte, _ := json.Marshal(updateRequest)

	req, rr := createRequestTest(http.MethodPut, fmt.Sprintf("/products/%s", "id"), string(dataByte))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 400, rr.Code)

}

func TestDelete(t *testing.T) {
	router := createServer()
	postRequest := handler.Request{
		Nombre:        "Post Product",
		Color:         "Pink",
		Precio:        200,
		Stock:         30,
		Codigo:        "3C3CFFJH0ET60XXXX",
		Publicado:     false,
		FechaCreacion: "26/08/2021",
	}

	dataByte, _ := json.Marshal(postRequest)

	req, rr := createRequestTest(http.MethodPost, "/products", string(dataByte))
	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	var response web.Response
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	prodByte, err := json.Marshal(response.Contenido)
	assert.Nil(t, err)
	var prod products.Product
	err = json.Unmarshal(prodByte, &prod)
	assert.Nil(t, err)
	assert.Equal(t, "Post Product", prod.Nombre)
	reqDel, rrDel := createRequestTest(http.MethodDelete, fmt.Sprintf("/products/%d", prod.ID), string(dataByte))
	router.ServeHTTP(rrDel, reqDel)

	assert.Equal(t, 200, rrDel.Code)
}
