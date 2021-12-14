/*
Se requiere probar la funcionalidad de “actualización de producto”, pudiendo reutilizar las funciones creadas en la clase.
Para lograrlo realizar los siguientes pasos:
	Dentro de la carpeta /test, crear un archivo products_test.go.
	Levantar el Servidor y definir la ruta para este test.
	Crear Request y Response apropiados.
	Solicitar al servidor que atienda al Request.
	Validar Response.
*/

/*
Se solicita probar la funcionalidad de “eliminar producto”, pudiendo reutilizar las funciones creadas en la clase.
Para lograrlo realizar los siguientes pasos:
	Dentro de la carpeta /test, crear un archivo products_test.go.
	Levantar el Servidor y definir la ruta para este test.
	Crear Request y Response apropiados.
	Solicitar al servidor que atienda al Request.
	Validar Response.
*/

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/TT/proyecto/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/12_testing3/TT/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/12_testing3/TT/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"Error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {

	tokenENV := os.Getenv("TOKEN")

	if tokenENV == "" {
		log.Fatal("Por favor seteá una token en .env")
	}

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "Token no enviado")
			return
		}

		if token != tokenENV {
			respondWithError(c, 401, "Token inválido")
			return
		}

		c.Next()
	}
}

func createServer() *gin.Engine {

	router := gin.Default()

	_ = os.Setenv("TOKEN", "123")

	db := store.New(store.FileType, "/Users/beconti/Desktop/meli_bootcamp2/12_testing3/TT/proyecto/tests/productosTest.json")

	repository := productos.NewRepository(db)
	service := productos.NewService(repository)
	controller := handler.NewProducto(service)

	router.Use(TokenAuthMiddleware())

	productos := router.Group("/productos")
	{
		productos.GET("/", controller.GetAll())
		productos.POST("/", controller.Store())
		productos.PUT("/:id", controller.Update())
		productos.DELETE("/:id", controller.Delete())
		productos.PATCH("/:id", controller.UpdateNombrePrecio())
	}

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123")

	return req, httptest.NewRecorder()
}

func Test_UpdateProducto(t *testing.T) {
	//Arrenge
	router := createServer()

	productToUpdate := productos.Producto{
		ID:            1,
		Nombre:        "Asdf",
		Color:         "Negro",
		Precio:        "$ 50",
		Stock:         2,
		Codigo:        "123456R",
		Publicado:     true,
		FechaCreacion: "01/06/1996",
	}

	url := fmt.Sprintf("/productos/%v", productToUpdate.ID)

	data, _ := json.Marshal(productToUpdate)

	request, response := createRequestTest(http.MethodPut, url, string(data))

	//Act
	router.ServeHTTP(response, request)

	//Assert
	assert.Equal(t, 200, response.Code)
}

func Test_DeleteProducto(t *testing.T) {
	//Arrenge
	router := createServer()

	productID := 7

	url := fmt.Sprintf("/productos/%v", productID)

	request, response := createRequestTest(http.MethodDelete, url, "")

	//Act
	router.ServeHTTP(response, request)

	//Assert
	assert.Equal(t, 200, response.Code)
}
