package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/afternoon/go-web/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/12_testing3/afternoon/go-web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/afternoon/go-web/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/afternoon/go-web/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "../cmd/server/productos.json")
	repo := internal.NewRepository(db)
	service := internal.NewService(repo)
	handler := handler.NewProducto(service)

	r := gin.Default()

	pr := r.Group("/productos")
	pr.PUT("/:id", handler.Update())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")
	return req, httptest.NewRecorder()
}

func Test_UpdateProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	sentProd := `{
		"nombre": "after Update",
		"color": "gris",
		"precio": 1,
		"stock": 12,
		"codigo": "sfsdf 444 3 www",
		"publicado": false,
		"fechaCreacion": "12/9/1999"
	  }`
	var objRes web.Response

	// crear Request del tipo Put y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/productos/1", sentProd)
	//guardo la respuesta en una estructura de tipo web.Response para hacer una evaluacion
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	assert.Nil(t, err)
	assert.Equal(t, 200, objRes.StatusCode, "deben ser iguales")
}

/* func Test_UpdateProduct_Fail(t *testing.T) { //falla porque se le envia un numero de id que no existe
	// crear el Server y definir las Rutas
	r := createServer()
	sentProd := `{
		"nombre": "after Update",
		"color": "gris",
		"precio": 1,
		"stock": 12,
		"codigo": "sfsdf 444 3 www",
		"publicado": false,
		"fechaCreacion": "12/9/1999"
	  }`
	//var objRes web.Response

	// crear Request del tipo Put y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/productos/9999", sentProd)
	//guardo la respuesta en una estructura de tipo web.Response para hacer una evaluacion
	//err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Code, "debe ser 400")

	/* assert.Nil(t, err)
	assert.Equal(t, 400, objRes, "deben ser iguales")
} */
