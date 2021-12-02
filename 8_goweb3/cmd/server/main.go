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

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.String(http.StatusUnauthorized, "Falta token")
		return false
	}
	secretToken := os.Getenv("TOKEN")
	if token != secretToken {
		ctx.String(http.StatusUnauthorized, "Token incorrecto")
		return false
	}

	return true
}

func MyMiddleware(c *gin.Context) {

	if !validarToken(c) {
		c.Abort()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}
	router := gin.Default()

	router.Use(MyMiddleware)

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
