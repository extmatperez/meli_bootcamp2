package main

import (
	"log"
	"net/http"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/soto_jose/8_goweb3/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/soto_jose/8_goweb3/pkg/store"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/soto_jose/8_goweb3/transactions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func validateAuthToken(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.String(http.StatusUnauthorized, "missing auth token")
		c.Abort()
		return
	}

	secretToken := os.Getenv("TOKEN")
	if token != secretToken {
		c.String(http.StatusUnauthorized, "invalid auth token")
		c.Abort()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}
	router := gin.Default()

	router.Use(validateAuthToken)

	db := store.New(store.FileType, "./transactions.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	controller := handler.NewTransaction(service)

	transactions := router.Group("/transactions")

	transactions.GET("", controller.GetAll())
	transactions.POST("", controller.Store())
	transactions.PUT("/:id", controller.Update())
	transactions.DELETE("/:id", controller.Delete())
	transactions.PATCH("/:id", controller.UpdateCodeAndAmount())

	router.Run()
}
