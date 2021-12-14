package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)



func createServer() *gin.Engine{
	db := store.New(store.FileType, "/transacciones.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	r := gin.Default()
	tr := r.Group("/transacciones")
	tr.POST("/", t.Store())
	tr.GET(("/"), t.GetAll())
	tr.PUT(":/id", t.Update())
	tr.PATCH(":/id", t.UpdateEmisor())
	tr.DELETE(":/id", t.Delete())

	return r
}


func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
    req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("token", "123456")

    return req, httptest.NewRecorder()
}


func TestUpdateTransaccion(t *testing.T){
	r := createServer()
	tran := transacciones.Transaccion{
		CodigoTransaccion: 486499,
		Moneda: "pesos",
		Monto: 84.16,
		Emisor: "LoloLC",
		Receptor: "Hauck-Carter",
		FechaTransaccion: "13/08/2021",
	  }
	  data, _ := json.Marshal(tran)
	req, rr := createRequestTest(http.MethodPut, "/transacciones/1", string(data))

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}



func TestDeleteTransaccion(t *testing.T){
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/transacciones/3", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}