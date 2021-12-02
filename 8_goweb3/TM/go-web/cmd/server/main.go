package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TM/go-web/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TM/go-web/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := productos.NewRepository()
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
