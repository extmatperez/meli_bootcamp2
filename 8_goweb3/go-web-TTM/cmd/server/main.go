package main

import (
	"fmt"
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTM/cmd/server/handler"
	product "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTM/internal/products"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTM/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./8_goweb3/go-web-TTMM/cmd/server/.env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("error al intentar cargar el archivo")
	}

	router := gin.Default()
	db := store.New(store.FileType, "./8_goweb3/go-web-TTM/cmd/server/products.json")
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/product/get", controller.GetAll())
	router.POST("/product/add", controller.Store())
	router.PUT("/product/put/:id", controller.Update())
	router.PATCH("/product/patch/:id", controller.UpdateNombre())
	router.PATCH("/product/patchprecio/:id", controller.UpdatePrecio())
	router.DELETE("/product/delete/:id", controller.Delete())
	router.Run()
}
