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

	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/12_testing3/go-web-TM/cmd/server/handler"
	product "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/12_testing3/go-web-TM/internal/products"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/12_testing3/go-web-TM/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/12_testing3/go-web-TM/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no hay variable de entorno de token cofigurada")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			respondWithError(ctx, 400, web.NewReponse(400, nil, "falta token"))
			//ctx.JSON(400, web.NewReponse(400, nil, "falta token"))
			return
		}
		tokenEnv := os.Getenv("TOKEN")
		if token != tokenEnv {
			respondWithError(ctx, 400, web.NewReponse(400, nil, "token incorrecto"))
			//ctx.JSON(400, web.NewReponse(404, nil, "token incorrecto"))
			return
		}
		ctx.Next()
	}
}
func CreateServe() *gin.Engine {

	err := godotenv.Load("../cmd/server/.env")
	if err != nil {
		log.Fatal("error al intentar cargar el archivo")
	}

	router := gin.Default()
	db := store.New(store.FileType, "./productsTest.json")
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	controller := handler.NewProduct(service)

	router.Use(TokenAuthMiddleware())
	router.GET("/product/get", controller.GetAll())
	router.POST("/product/add", controller.Store())
	router.PUT("/product/put/:id", controller.Update())
	router.PATCH("/product/patch/:id", controller.UpdateNombre())
	router.PATCH("/product/patchprecio/:id", controller.UpdatePrecio())
	router.DELETE("/product/delete/:id", controller.Delete())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123")

	return req, httptest.NewRecorder()
}
func Test_GetProducts(t *testing.T) {
	router := CreateServe()

	req, rr := createRequestTest(http.MethodGet, "/product/get", "")

	// indicar al servidor que pueda atender la solicitud
	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	var respuesta web.Response
	fmt.Println(respuesta)
	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	fmt.Println(respuesta)
	assert.Nil(t, err)
	assert.Equal(t, 200, respuesta.Code)

	//assert.True(t, len(respuesta.Data) > 0)

}

func Test_StoreProducts(t *testing.T) {
	router := CreateServe()

	productoNuevo := product.Product{
		Nombre:        "producto",
		Color:         "rojo",
		Precio:        20,
		Stock:         "alguno",
		Codigo:        "SADFHJK9",
		Publicado:     true,
		FechaCreacion: "01/12/2021",
	}
	dataNueva, _ := json.Marshal(productoNuevo)
	fmt.Println(string(dataNueva))
	req, rr := createRequestTest(http.MethodPost, "/product/add", string(dataNueva))

	// indicar al servidor que pueda atender la solicitud
	router.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	var res product.Product

	//assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &res)

	fmt.Println(res)
	assert.Nil(t, err)

	assert.Equal(t, productoNuevo.Nombre, res.Nombre)

}
