package main

import (
	handler "github.com/extmatperez/meli_bootcamp2/8_goweb3/morning/go-web/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/8_goweb3/morning/go-web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/morning/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db := store.New(store.FileType, "./productos.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	controller := handler.NewProducto(service)

	router := gin.Default()

	router.GET("/productos/", controller.GetAll())
	router.POST("productos", controller.Store())
	router.PUT("productos/:id", controller.Update())
	router.PATCH("productos/:id", controller.UpdateName())
	router.DELETE("productos/:id", controller.Delete())

	router.Run(":8080")
}
