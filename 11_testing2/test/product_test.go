package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	handlers "github.com/extmatperez/meli_bootcamp2/11_testing2/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/11_testing2/internal/products"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/pkg/web"
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
	router := gin.Default()
	_ = os.Setenv("TOKEN", "TOKEN-PRODUCTS")

	router.Use(TokenAuthMiddleware())

	dbStore := store.New(store.FileType, "./products-test.json")
	productsRepository := products.NewRepository(dbStore)
	productsService := products.NewService(productsRepository)
	productsController := handlers.NewProduct(productsService)

	products := router.Group("/products")
	{
		products.GET("/", productsController.GetAll())
		products.GET("/:id", productsController.FindById())
		products.POST("/", productsController.Store())
		products.PUT("/:id", productsController.Update())
		products.DELETE("/:id", productsController.Delete())
		products.PATCH("/:id", productsController.UpdateNameAndPrice())
	}

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "TOKEN-PRODUCTS")

	return req, httptest.NewRecorder()
}

func Test_GetAllProducts(t *testing.T) {
	router := createServer()
	req, res := createRequestTest(http.MethodGet, "/products/", "")

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)

	var response web.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)

	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}

func Test_StoreProduct(t *testing.T) {
	router := createServer()

	newProduct := products.Product{
		Name:       "Boligoma",
		Color:      "Transparente",
		Price:      200.50,
		Stock:      10,
		Code:       "#2000030",
		Published:  true,
		Created_at: "14/12/2021",
	}

	newProductBytes, _ := json.Marshal(newProduct)
	req, res := createRequestTest(http.MethodPost, "/products/", string(newProductBytes))

	router.ServeHTTP(res, req)

	assert.Equal(t, 201, res.Code)

	var response web.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)

	assert.Equal(t, 201, response.Code)
	assert.Nil(t, err)
}

func Test_UpdateProduct(t *testing.T) {
	router := createServer()

	updateProduct := `
		{
			"name": "boligoma act",
			"color":      "Transparente",
			"price":      200.50,
			"stock":      10,
			"code":       "#2000030",
			"published":  true,
			"created_at": "14/12/2021"
		}
	`

	req, res := createRequestTest(http.MethodPut, "/products/1", string(updateProduct))

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)

	var response web.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)

	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}

func Test_DeleteProduct(t *testing.T) {
	router := createServer()

	req, res := createRequestTest(http.MethodDelete, "/products/1", "")

	router.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)

	var response web.Response
	err := json.Unmarshal(res.Body.Bytes(), &response)

	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}
