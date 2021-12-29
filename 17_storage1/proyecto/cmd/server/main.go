package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/docs"
	transacciones "github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

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

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "../../internal/transacciones/transacciones.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	r := gin.Default()
	//r.Use(TokenAuthMiddleware())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := r.Group("/transacciones")
	tr.POST("/add", TokenAuthMiddleware(), t.Store())
	tr.POST("/load", TokenAuthMiddleware(), t.Load())
	tr.GET("/get", t.GetAll())
	tr.GET("/find/:id", t.FindById())
	tr.GET("/filter", t.FilterBy())
	tr.PUT("/update/:id", TokenAuthMiddleware(), t.Update())
	tr.PATCH("/cod/:id", TokenAuthMiddleware(), t.UpdateCod())
	tr.PATCH("/mon/:id", TokenAuthMiddleware(), t.UpdateMon())
	tr.DELETE("/del/:id", TokenAuthMiddleware(), t.Delete())
	//tr.DELETE("/delAll",TokenAuthMiddleware(), t.Delete())

	err := r.Run()

	if err != nil {
		panic(err)
	}

}
