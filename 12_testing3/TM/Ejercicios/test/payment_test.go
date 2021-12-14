package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/12_testing3/TM/Ejercicios/cmd/server/handler"
	payments "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/12_testing3/TM/Ejercicios/internal/payments"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/12_testing3/TM/Ejercicios/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/12_testing3/TM/Ejercicios/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Hacemos el control del token aqui tambien.
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	// Toma el token del .env
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Por favor, setear la variable de entorno TOKEN.")
	}

	return func(c *gin.Context) {
		// Toma el token del header en el postman.
		token := c.GetHeader("TOKEN")

		if token == "" {
			respondWithError(c, 401, "API Token requerido.")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "API Token inválido.")
			return
		}

		c.Next()
	}
}

func createServer() *gin.Engine {
	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./payments.json")
	repository := payments.NewRepository(db)
	service := payments.NewService(repository)
	controller := handler.NewPayment(service)

	router.Use(TokenAuthMiddleware())

	payments := router.Group("/payments")
	{
		payments.GET("/get", controller.GetAll())
		payments.POST("/", controller.Store())
		payments.PUT("/:id", controller.Update())
		payments.PATCH("/code/:id", controller.UpdateCodigo())
		payments.PATCH("/amount/:id", controller.UpdateMonto())
		payments.DELETE("/:id", controller.Delete())
	}

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()
}

// Aca se corren los tests funcionales, obteniendo cada uno de los endpoints para probar.
func Test_GetPayment(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/payments/get", "")

	// A partir de aca se van a guardar los datos en el response (rr)
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Code)
	assert.Nil(t, err)
}

func Test_StorePayment(t *testing.T) {
	router := createServer()

	newPayment := payments.Payment{Codigo: "0000004", Moneda: "ARS", Monto: 856.34, Emisor: "Rodrigo Vega Gimenez", Receptor: "Nicolas Valentinuzzi", Fecha: "2021-12-11"}

	newData, _ := json.Marshal(newPayment)
	req, rr := createRequestTest(http.MethodPost, "/payments/", string(newData))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 201, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, 201, response.Code)
	assert.Nil(t, err)
}

func Test_UpdatePayment(t *testing.T) {
	router := createServer()

	updatedPayment := payments.Payment{Codigo: "0000004.1", Moneda: "ARS", Monto: 856.34, Emisor: "Rodrigo Vega", Receptor: "Nicolas Valentinuzzi", Fecha: "2021-12-11"}

	newData, _ := json.Marshal(updatedPayment)
	req, rr := createRequestTest(http.MethodPut, "/payments/1", string(newData))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}

func Test_UpdatePaymentCode(t *testing.T) {
	router := createServer()

	updatedPayment := payments.Payment{Codigo: "0000005"}

	newData, _ := json.Marshal(updatedPayment)
	req, rr := createRequestTest(http.MethodPatch, "/payments/code/1", string(newData))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}

func Test_UpdatePaymentAmount(t *testing.T) {
	router := createServer()

	updatedPayment := payments.Payment{Monto: 2356.34}

	newData, _ := json.Marshal(updatedPayment)
	req, rr := createRequestTest(http.MethodPatch, "/payments/amount/1", string(newData))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}

func Test_DeletePayment(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodDelete, "/payments/3", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var response web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &response)

	assert.Equal(t, 200, response.Code)
	assert.Nil(t, err)
}
