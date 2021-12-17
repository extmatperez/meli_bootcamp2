package test

import (
	"bytes"
	"encoding/json"
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

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
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
	pr := router.Group("/transactions")
	pr.PATCH("/", t.UpdateAmount())
	pr.GET("/", t.GetAll())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "dig.123")

	return req, httptest.NewRecorder()
}

func Test_GetAll(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/transactions", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}

//Update
func Test_UpdateTransaction(t *testing.T) {
	router := createServer()

	//updateTrans := transaction.Transaction{Amount: 1.00}
	req, rr := createRequestTest(http.MethodPatch, "/transactions/1/200.00", "")
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Code)
	assert.Nil(t, err)
}

//Delete
