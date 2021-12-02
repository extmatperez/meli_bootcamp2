package main

import (
	"log"

	productos "github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/go-web/clase3/goModularizadoEnCapas/Internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/go-web/clase3/goModularizadoEnCapas/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/zamora_damian/go-web/clase3/goModularizadoEnCapas/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.New(store.FileType, "./products.json")

	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	controller := handler.NewPersona(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/modificar/:id", controller.Update())
	//router.PATCH("/modificarNombre/:id", controller.UpdateName())
	//router.DELETE("/delete/:id", controller.delete())
	router.Run()
}
