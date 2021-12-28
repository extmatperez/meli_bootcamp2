package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/17_storage1/Go_Web/cmd/server/handler"
	usuarios "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/17_storage1/Go_Web/internal/usuarios"
	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/17_storage1/Go_Web/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/8_goweb3/Go_Web/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	router := gin.Default()

	db := store.New(store.FileType, "./usuariosSalidaTest.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	controller := handler.NewUsuario(service)

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("usuarios/update", controller.Update())
	router.DELETE("usuarios/delete/:id", controller.Delete())
	router.PATCH("usuarios/patch/:id", controller.EditarNombreEdad())

	return router
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json ")
	//req.Header.Add("token", "12345")
	return req, httptest.NewRecorder()
}

func Test_GetUser_OK(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/usuarios/get", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)

	assert.Equal(t, 200, respuesta.Code)

	assert.Nil(t, err)

}

func Test_StoreUser(t *testing.T) {
	router := createServer()

	nuevoUsuario := usuarios.Usuario{Nombre: "Juan", Apellido: "Pascal", Email: "email", Edad: 20, Altura: 172, Activo: true, FechaCreacion: "2019"}
	dataNueva, _ := json.Marshal(nuevoUsuario)
	req, rr := createRequestTest(http.MethodPost, "/usuarios/add", string(dataNueva))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta usuarios.Usuario

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)

	assert.Equal(t, "Juan", respuesta.Nombre)

	assert.Nil(t, err)

}

func Test_UpdateUser(t *testing.T) {
	router := createServer()

	nuevoUsuario := usuarios.Usuario{ID: 10, Nombre: "Juan", Apellido: "Pascal", Email: "email", Edad: 20, Altura: 172, Activo: true, FechaCreacion: "2019"}
	dataNueva, _ := json.Marshal(nuevoUsuario)

	req, rr := createRequestTest(http.MethodPut, "/usuarios/update", string(dataNueva))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	fmt.Println(respuesta)
	assert.Equal(t, 200, respuesta.Code)

	assert.Nil(t, err)
}

func Test_DeleteUser(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodDelete, "/usuarios/delete/8", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	fmt.Println(respuesta)
	assert.Equal(t, 200, respuesta.Code)

	assert.Nil(t, err)
}
