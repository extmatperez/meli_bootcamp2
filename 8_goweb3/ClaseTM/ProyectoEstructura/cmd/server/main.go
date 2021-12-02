package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/8_goweb3/ClaseTM/ProyectoEstructura/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/8_goweb3/ClaseTM/ProyectoEstructura/internal/producto"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/8_goweb3/ClaseTM/ProyectoEstructura/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.NewStore(store.FileType, "./dbProductos.json")
	repo := producto.NewRepository(db)
	service := producto.NewService(repo)
	controller := handler.NewProductoController(service)

	groupProducts := router.Group("api/productos")
	{
		groupProducts.GET("/", controller.GetAll())
		groupProducts.POST("/", controller.Store())
		groupProducts.PUT("/:id", controller.Update())
		groupProducts.DELETE("/:id", controller.Detele())
		groupProducts.PATCH("/:id", controller.UpdateNameAndPrice())

	}

	router.Run()
}
