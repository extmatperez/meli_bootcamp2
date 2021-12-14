package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/12_testing3/afternoon/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/12_testing3/afternoon/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/12_testing3/afternoon/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/12_testing3/afternoon/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.NewStore(store.FileType, "./../cmd/server/users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	hand := handler.NewUser(service)

	r := gin.Default()
	pr := r.Group("/users")
	pr.PUT("/:id", hand.Update())
	pr.DELETE("/:id", hand.Delete())
	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestUpdateUserOK(t *testing.T) {
	userJson := `{
		"name": "Juan",
		"last_name": "Perez",
		"email": "juan@perez.com",
		"age": 45,
		"height": 1.75,
		"active" : true,
		"created": "01/12/2021"
	}`

	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/users/2", userJson)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	user := respuesta.Data.(map[string]interface{})

	assert.Nil(t, err)
	assert.Equal(t, "Juan", user["name"])
	assert.Equal(t, 2.0, user["id"])

}

func TestDeleteUserOK(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/users/2", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)

	assert.Nil(t, err)
	assert.Equal(t, 200, respuesta.Code)
	assert.Equal(t, "The user 2 has been deleted", respuesta.Data)
}
