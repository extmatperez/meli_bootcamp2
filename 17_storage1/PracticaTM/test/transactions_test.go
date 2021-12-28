package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTM/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTM/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/17_storage1/PracticaTM/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	// _ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./transaccionesSalida.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	p := handler.NewTransaccion(service)
	r := gin.Default()
	pr := r.Group("/transactions")
	pr.POST("/", p.Store())
	pr.PUT("/", p.Update())
	pr.GET("/", p.GetAll())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")
	return req, httptest.NewRecorder()
}

func Test_UpdateTransactions_OK(t *testing.T) {
	var objRes []transacciones.Transaccion
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo Post y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/transactions/", `{
        "id": 5, "cod_transaccion": "987ZYX","moneda": "pesos","monto": 999.987,"emisor": "Mariano","receptor": "Martin","fecha_trans": "21/01/2021"
    }`)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)
}

// func Test_SaveProduct_OK(t *testing.T) {
// 	// crear el Server y definir las Rutas
// 	r := createServer()
// 	// crear Request del tipo POST y Response para obtener el resultado
// 	req, rr := createRequestTest(http.MethodPost, "/transactions/", `{
//         "cod_transaccion": "987ZYX","moneda": "pesos uruguayos","emisor": "Mariano","receptor": "Martin","fecha_trans": "21-01-2021"
//     }`)
// 	// indicar al servidor que pueda atender la solicitud
// 	r.ServeHTTP(rr, req)
// 	assert.Equal(t, 200, rr.Code)
// }
