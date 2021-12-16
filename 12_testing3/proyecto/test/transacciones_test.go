package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/proyecto/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/proyecto/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	transacciones "github.com/extmatperez/meli_bootcamp2/12_testing3/proyecto/internal/transacciones"
)

func createServer() *gin.Engine {
	r := gin.Default()

	db := store.New(store.FileType, "./transaccionTest.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	tr := r.Group("/transacciones")
	tr.POST("/add" /* TokenAuthMiddleware(), */, t.Store())
	tr.POST("/load" /* TokenAuthMiddleware(), */, t.Load())
	tr.GET("/get", t.GetAll())
	tr.GET("/find/:id", t.FindById())
	tr.GET("/filter", t.FilterBy())
	tr.PUT("/update/:id" /* TokenAuthMiddleware(), */, t.Update())
	tr.PATCH("/cod/:id" /* TokenAuthMiddleware(), */, t.UpdateCod())
	tr.PATCH("/mon/:id" /* TokenAuthMiddleware(), */, t.UpdateMon())
	tr.DELETE("/del/:id" /* TokenAuthMiddleware(), */, t.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetTransacciones(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/transacciones/get", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Code)
	assert.Nil(t, err)
}

func Test_StorePersonas(t *testing.T) {
	router := createServer()

	nuevaTrans := transacciones.Transaccion{Moneda: "pesoArg", Monto: 3014.4, Emisor: "roro", Receptor: "maslata"}

	dataNueva, _ := json.Marshal(nuevaTrans)
	req, rr := createRequestTest(http.MethodPost, "/transacciones/add", string(dataNueva))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	trans := respuesta.Data.(map[string]interface{})
	assert.Equal(t, "roro", trans["emisor"])
	assert.Nil(t, err)
}

func Test_UpdatePersonas(t *testing.T) {
	router := createServer()

	nuevaTrans := transacciones.Transaccion{Moneda: "pesrg", Monto: 301.4, Emisor: "roasro", Receptor: "masdfglata"}

	dataNueva, _ := json.Marshal(nuevaTrans)
	req, rr := createRequestTest(http.MethodPut, "/transacciones/update/1", string(dataNueva))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	//var respuesta transacciones.Transaccion
	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	trans := respuesta.Data.(map[string]interface{})
	//assert.Equal(t, "roasro", respuesta.Emisor)
	assert.Equal(t, "roasro", trans["emisor"])

	assert.Nil(t, err)

}

func Test_DeletePersonas(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodDelete, "/transacciones/del/102", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 201, rr.Code)

}
