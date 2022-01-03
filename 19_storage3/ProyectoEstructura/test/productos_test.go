package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/internal/producto"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/pkg/store"

	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {
	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")
	db := store.NewStore(store.FileType, "./personasSalidaTest.json")
	repo := producto.NewRepository(db)
	service := producto.NewService(repo)
	controller := handler.NewProductoController(service)

	// router.Use(TokenAuthMiddleware())
	// data := respuesta.Contenido.(map[string]interface{})
	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}
