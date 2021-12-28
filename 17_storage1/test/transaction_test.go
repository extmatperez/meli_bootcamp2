package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_afternoon/cmd/server/handler"
	transaction "github.com/extmatperez/meli_bootcamp2/12_testing3/turn_afternoon/internal/transaction"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_afternoon/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_afternoon/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = createServer()
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		c.Next()
	}
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "dig.123")
	db := store.New(store.FileType, "./transaction.json")
	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	t := handler.NewTransaction(service)
	router := gin.Default()

	transactionURL := router.Group("/transactions")
	transactionURL.PATCH("/:id/:amount", t.UpdateAmount())
	transactionURL.GET("/", t.GetAll())
	transactionURL.GET("/:id", t.GetByID())
	transactionURL.GET("/receivers/:receiver", t.GetByReceiver())
	transactionURL.POST("/", t.Store())
	transactionURL.PUT("/:id", t.UpdateTransaction())
	transactionURL.DELETE("/:id", t.DeleteTransaction())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "dig.123")

	return req, httptest.NewRecorder()
}

func Test_GetAll(t *testing.T) {

	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}

//Store
func Test_StoreTransactions(t *testing.T) {

	newTransaction := transaction.Transaction{
		ID:              13,
		TransactionCode: "1234-2124",
		Currency:        "$",
		Amount:          100.00,
		Receiver:        "receiver",
		Sender:          "sender",
		TransactionDate: "12/12/2021",
	}

	newDataTrans, _ := json.Marshal(newTransaction)
	fmt.Println(string(newDataTrans))
	req, rr := createRequestTest(http.MethodPost, "/transactions/", string(newDataTrans))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	fmt.Println("Body response: ", rr.Body)
	var response transaction.Transaction

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	fmt.Println("response: ", response)

	assert.Nil(t, err)

}

//Delete
func Test_DeleteTransaction(t *testing.T) {

	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	fmt.Println("response delete: ", response.Transaction)

	assert.Nil(t, err)

	trans := []transaction.Transaction{}
	bodyBytes, _ := json.Marshal(response.Transaction)
	json.Unmarshal(bodyBytes, &trans)
	fmt.Println(trans)

	delete := fmt.Sprintf("/transactions/%d", trans[len(trans)-1].ID)
	req, rr = createRequestTest(http.MethodDelete, delete, "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}

//Update
func Test_UpdateTransaction(t *testing.T) {

	//updateTrans := transaction.Transaction{Amount: 1.00}
	req, rr := createRequestTest(http.MethodPatch, "/transactions/1/200.00", "")
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
