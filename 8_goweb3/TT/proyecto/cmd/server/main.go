package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "../../internal/transacciones/transacciones.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	r := gin.Default()
	//r.Use(TokenAuthMiddleware())
	tr := r.Group("/transacciones")
	tr.POST("/add", TokenAuthMiddleware(), t.Store())
	tr.GET("/get", t.GetAll())
	tr.GET("/load", TokenAuthMiddleware(), t.Load())
	tr.GET("/find/:id", t.FindById())
	tr.GET("/filter", t.FilterBy())
	tr.PUT("/update/:id", TokenAuthMiddleware(), t.Update())
	tr.PATCH("/cod/:id", TokenAuthMiddleware(), t.UpdateCod())
	tr.PATCH("/mon/:id", TokenAuthMiddleware(), t.UpdateMon())
	tr.DELETE("/del/:id", TokenAuthMiddleware(), t.Delete())
	//tr.DELETE("/delAll",TokenAuthMiddleware(), t.Delete())

	r.Run()
}
