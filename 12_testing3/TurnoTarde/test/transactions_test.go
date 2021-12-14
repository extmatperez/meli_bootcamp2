package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/12_testing3/TurnoTarde/cmd/server/handler"
	tran "github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/12_testing3/TurnoTarde/internal/transaccion"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/12_testing3/TurnoTarde/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/12_testing3/TurnoTarde/pkg/store/web"
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

func Test_Update_OK(t *testing.T) {
	// crear el Server y definir las Rutas
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

	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, url, bodyByte)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	var bodyRecived web.Response
	var transactionRecived tran.Transaction
	json.Unmarshal(rr.Body.Bytes(), &bodyRecived)

	byt, _ := json.Marshal(bodyRecived.Content)
	json.Unmarshal(byt, &transactionRecived)
	fmt.Println(rr)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, transaction.Codigo, transactionRecived.Codigo)
	assert.Equal(t, transaction.Emisor, transactionRecived.Emisor)
	assert.Equal(t, transaction.Fecha, transactionRecived.Fecha)
	assert.Equal(t, id, transactionRecived.ID)
	assert.Equal(t, transaction.Moneda, transactionRecived.Moneda)
	assert.Equal(t, transaction.Monto, transactionRecived.Monto)
}
