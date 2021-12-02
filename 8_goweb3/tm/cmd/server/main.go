package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/8_goweb3/tm/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/8_goweb3/tm/internal/products"
	store "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/8_goweb3/tm/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	router := gin.Default()

	db := store.NewStore("file", "../prueba.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("getAll", controller.GetAll())
	router.POST("store", controller.Store())
	router.PUT("update", controller.Update())

	router.Run()
}
