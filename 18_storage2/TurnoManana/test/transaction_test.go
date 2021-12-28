package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoManana/cmd/server/handler"
	tran "github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoManana/internal/transaccion"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoManana/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoManana/pkg/store/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	server := gin.Default()
	os.Setenv("TOKEN", "123456")
	//inicializaciones
	db := store.New(store.FileType, "/Users/frapalacio/Documents/Repositorios/meli_bootcamp2/12_testing3/TurnoTarde/cmd/transactions.json")
	repo := tran.NewRepository(db)
	service := tran.NewService(repo)
	controller := handler.NewTransaction(service)

	transaction := server.Group("/transactions")

	{
		//get
		transaction.GET("/", controller.GetAll())
		transaction.GET("/:id", controller.GetTransactionById())
		transaction.GET("/filtros", controller.GetTransactionsExlusive())

		//post
		transaction.POST("/", controller.Store())

		//put
		transaction.PUT("/:id", controller.Update())

		//patch
		transaction.PATCH("/:id", controller.UpdateCodigoAndMonto())

		//delete
		transaction.DELETE("/:id", controller.Delete())
	}
	return server
}

func createRequestTest(method string, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}

func TestUpdateSucces(t *testing.T) {
	r := createServer()
	id := 2
	transaction := handler.Request{
		Codigo:   "Cod123456",
		Moneda:   "Ars",
		Monto:    "5000",
		Emisor:   "PANCHO",
		Receptor: "AGUSTO",
		Fecha:    "14/12/21",
	}

	bodyByte, _ := json.Marshal(transaction)
	url := fmt.Sprintf("/transactions/%v", id)


	req, rr := createRequestTest(http.MethodPut, url, bodyByte)


	r.ServeHTTP(rr, req)

	var bodyRecived web.Response
	var transactionRecived tran.Transaction
	
	json.Unmarshal(rr.Body.Bytes(), &bodyRecived)

	byt, _ := json.Marshal(bodyRecived.Content)
	json.Unmarshal(byt, &transactionRecived)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, transaction.Codigo, transactionRecived.Codigo)
	assert.Equal(t, transaction.Emisor, transactionRecived.Emisor)
	assert.Equal(t, transaction.Fecha, transactionRecived.Fecha)
	assert.Equal(t, id, transactionRecived.ID)
	assert.Equal(t, transaction.Moneda, transactionRecived.Moneda)
	assert.Equal(t, transaction.Monto, transactionRecived.Monto)
}




func TestDeleteSucces(t *testing.T) {
	r := createServer()
	transaction := handler.Request{
		Codigo:   "Cod123456",
		Moneda:   "Ars",
		Monto:    "5000",
		Emisor:   "PANCHO",
		Receptor: "AGUSTO",
		Fecha:    "14/12/21",
	}

	bodyByte, _ := json.Marshal(transaction)
	url := fmt.Sprintf("/transactions/")

	req, rr := createRequestTest(http.MethodPost, url, bodyByte)

	r.ServeHTTP(rr, req)

	var bodyRecived web.Response
	var transactionRecived tran.Transaction
	
	json.Unmarshal(rr.Body.Bytes(), &bodyRecived)

	byt, _ := json.Marshal(bodyRecived.Content)
	json.Unmarshal(byt, &transactionRecived)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, transaction.Codigo, transactionRecived.Codigo)
	assert.Equal(t, transaction.Emisor, transactionRecived.Emisor)
	assert.Equal(t, transaction.Fecha, transactionRecived.Fecha)
	assert.Equal(t, transaction.Moneda, transactionRecived.Moneda)
	assert.Equal(t, transaction.Monto, transactionRecived.Monto)


	url = fmt.Sprintf("/transactions/%d",transactionRecived.ID)

	req, rr = createRequestTest(http.MethodDelete, url, bodyByte)
	
	r.ServeHTTP(rr, req)
	json.Unmarshal(rr.Body.Bytes(), &bodyRecived)
	assert.Equal(t, 200, rr.Code)



}
