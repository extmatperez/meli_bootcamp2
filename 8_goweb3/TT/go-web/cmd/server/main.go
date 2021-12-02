package main

import (
	"fmt"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error al intentar cargar el archivo.env")
	}
	file := store.New(store.FileType, "./productos.json")
	repo := productos.NewRepository(file)
	service := productos.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/productos")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/", p.Edit())
	pr.PATCH("/:id", p.Change())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
