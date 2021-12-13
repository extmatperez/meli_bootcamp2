package main

import (
	"log"

	handler "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error al intentar cargar el archivo .env")
	}
	db := store.New(store.FileType, "/transacciones.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	r := gin.Default()
	tr := r.Group("/transacciones")
	tr.POST("/", t.Store())
	tr.GET(("/"), t.GetAll())
	tr.PUT(":/id", t.Update())
	tr.PATCH(":/id", t.UpdateEmisor())
	tr.DELETE(":/id", t.Delete())

	r.Run()

}
