package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/extmatperez/meli_bootcamp2/tree/soto_jose/11_testing2/TT/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/soto_jose/11_testing2/TT/server/handler"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/soto_jose/11_testing2/TT/transactions"
	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {
	router := gin.Default()
	db := store.New(store.FileType, "./personasSalidaTest.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	controller := handler.NewTransaction(service)

	transactionsGroup := router.Group("/transactions")

	transactionsGroup.GET("", controller.GetAll())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}
