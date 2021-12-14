package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/12_testing3/go-web/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/12_testing3/go-web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/12_testing3/go-web/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/12_testing3/go-web/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./productosTest.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/", controller.GetAll())
	router.POST("/", controller.Store())
	router.PUT("/:id", controller.Edit())
	router.PATCH(":id", controller.Change())
	router.DELETE(":id", controller.Delete())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetProductos(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest("GET", "/", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Codigo)
	assert.Nil(t, err)
}

func Test_PutProductos(t *testing.T) {
	router := createServer()

	productoEditado := productos.Producto{Nombre: "ProductoTesting", Color: "Rojo", Precio: "$45.00", Stock: 24, Codigo: "abcd", Publicado: true, Creado: "25/12/2003"}
	bodyString, _ := json.Marshal(productoEditado)
	req, rr := createRequestTest("PUT", "/1", string(bodyString))

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Codigo)
	assert.Nil(t, err)
}

func Test_DeleteProductos(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest("DELETE", "/3", "")

	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Codigo)
	assert.Nil(t, err)
}
