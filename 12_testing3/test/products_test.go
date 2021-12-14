package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/12_testing3/internal/products"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	router := gin.Default()

	db := store.NewStore("file", "productsTest.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/products", controller.GetAll())
	router.POST("/products/addproduct", controller.AddProduct())
	router.PUT("/products/updateproduct/:id", controller.UpdateProduct())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()

}

func Test_GetProduct_OK(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/products", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	fmt.Print("Esto es rr: ", rr)
	fmt.Print("Esto es Body: ", rr.Body)
	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)

	assert.Nil(t, err)
	assert.Equal(t, 200, respuesta.Code)

}

func Test_AddProduct_OK(t *testing.T) {
	router := createServer()

	nuevoProducto := products.Product{Name: "Fruit", Color: "red", Price: 23, Stock: 30, Code: 656, Published: "5/17/2021", Created: "8/28/2021"}

	productByte, _ := json.Marshal(nuevoProducto)

	req, rr := createRequestTest(http.MethodPost, "/products/addproduct", string(productByte))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)

	assert.Nil(t, err)
	assert.Equal(t, 200, respuesta.Code)

}
