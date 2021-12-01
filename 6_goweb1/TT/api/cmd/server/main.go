package main

import (
	producto "github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/internal/producto"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repository := producto.NewRepository()
	service := producto.NewService(repository)
	controller := handler.newProduct(service)
}
