package test

/* import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/cmd/server/handler"
	personas "github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/personas"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/store"
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
	r := gin.Default()

	db := store.New(store.FileType, "./personasSalidaTest.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	r.Use(TokenAuthMiddleware())

	personasEP := r.Group("/personas")
	{
		personasEP.GET("/", controller.GetAll())
		personasEP.POST("/add", controller.Store())
		personasEP.PUT("/update/:id", controller.Update())
		personasEP.PATCH("/updateParcial/:id", controller.UpdateNombre())
		personasEP.DELETE("/delete/:id", controller.Delete())
	}

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *http.Response) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req
} */