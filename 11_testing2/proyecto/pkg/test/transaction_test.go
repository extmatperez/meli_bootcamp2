package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/internal/transactions"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	db := store.New(store.FileType, "./transactionsTest.json")
	repo := internal.NewRepository(db)
	service := internal.NewService(repo)
	controller := handler.NewTransaction(service)
	router := gin.Default()

	routerTransactions := router.Group("/transactions")
	{
		routerTransactions.GET("/", controller.GetAll())
		routerTransactions.GET("/:id", controller.GetTransactionByID())
		routerTransactions.POST("/", controller.Store())
		routerTransactions.PUT("/:id", controller.Update())
		routerTransactions.DELETE("/:id", controller.Delete())
		routerTransactions.PATCH("/:id", controller.UpdateCodigoYMonto())
	}
	return router
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")
	return req, httptest.NewRecorder()
}

func Test_Update_OK(t *testing.T) {
	r := createServer()
	url := fmt.Sprintf("/transactions/%d", 7)
	body := `{
		"codigo_de_transaccion": "asd12321",
		"moneda": "USD",
		"monto": 23400.85,
		"emisor": "sebac2",
		"receptor": "rec",
		"fecha_de_transaccion": "10/18/2021"
	}`
	req, rr := createRequestTest(http.MethodPut, url, body)
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	var respuesta web.Response
	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Nil(t, err)
	assert.NotNil(t, respuesta)
	assert.Equal(t, 200, respuesta.Code)

	assert.Equal(t, "asd12321", respuesta.Data.(map[string]interface{})["codigo_de_transaccion"])
}

func Test_Delete_OK(t *testing.T) {
	r := createServer()
	url := fmt.Sprintf("/transactions/%d", 8)

	req, rr := createRequestTest(http.MethodDelete, url, "")
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	var respuesta web.Response
	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Nil(t, err)
	assert.NotNil(t, respuesta)
	assert.Equal(t, 200, respuesta.Code)
	expected := fmt.Sprintf("transaction %d deleted", 8)
	assert.Equal(t, expected, respuesta.Data)
}
