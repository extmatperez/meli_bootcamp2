package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/12_testing3/ejTM/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func CreateServer() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldnt load the environment")
	}
	router := gin.Default()
	storage := store.NewStore("file", "productosTest.json")
	repository := internal.NewRepository(storage)
	service := internal.NewService(repository)
	producto := handler.NewProducto(service)

	router.GET("/producto", producto.GetAll())
	router.POST("/producto", producto.Store())
	router.GET("/producto/:id", producto.GetProductById())
	router.PUT("/producto/:id", producto.Update())
	router.DELETE("/producto/:id", producto.Delete())
	router.PATCH("/producto/:id", producto.UpdateNombrePrecio())

	return router
}

func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")
	return req, httptest.NewRecorder()

}

func Test_GetAllProductos(t *testing.T) {
	router := CreateServer()

	req, response := CreateRequestTest(http.MethodGet, "/producto", "")
	router.ServeHTTP(response, req)

	var respuesta web.Response

	json.Unmarshal(response.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Code)

}

func Test_StoreProductos(t *testing.T) {
	router := CreateServer()
	productoToPost := internal.Producto{Nombre: "nuevoProducto", Precio: 38.9, Color: "naranja"}
	byteBody, _ := json.Marshal(productoToPost)
	req, response := CreateRequestTest(http.MethodPost, "/producto", string(byteBody))
	router.ServeHTTP(response, req)

	var respuesta web.Response

	json.Unmarshal(response.Body.Bytes(), &respuesta)
	data := respuesta.Contenido.(map[string]interface{})
	assert.Equal(t, 200, respuesta.Code, "la operacion debe ser exitosa")
	assert.NotNil(t, respuesta.Contenido, "en el contenido tiene que venir algo")
	assert.Equal(t, productoToPost.Nombre, data["nombre"], "se debe crear con el nombre que se le pasa por body")

}

func Test_DeleteProducto(t *testing.T) {
	router := CreateServer()
	req, response := CreateRequestTest(http.MethodGet, "/producto", "")
	router.ServeHTTP(response, req)

	var respuesta web.Response

	json.Unmarshal(response.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Code, "la operacion debe ser exitosa")

	data := respuesta.Contenido.([]interface{})
	if len(data) != 0 {
		lastId := data[len(data)-1].(map[string]interface{})["id"]
		intId := int(lastId.(float64))
		stringID := strconv.Itoa(intId)
		// agarramos el id dle ultimo producto en la lista
		req, res := CreateRequestTest(http.MethodDelete, "/producto/"+stringID, "")
		router.ServeHTTP(response, req)

		var respuestaDelete web.Response

		json.Unmarshal(res.Body.Bytes(), &respuestaDelete)
		fmt.Println("RESPUESTA", res.Body)
		assert.Equal(t, 200, respuestaDelete.Code)
	}
	// assert.NotNil(t, respuesta.Contenido, "en el contenido tiene que venir algo")

}
