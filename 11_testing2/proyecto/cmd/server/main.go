package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/docs"
	producto "github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 0.1
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./productos.json")
	repo := producto.NewRepository(db)
	service := producto.NewService(repo)
	controller := handler.NewProducto(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/productos", controller.GetAll())
	router.POST("/addProductos", controller.Store())
	router.PUT("/modify/:id", controller.Modify())
	router.PATCH("/modifyNaPr/:id", controller.ModifyNamePrice())
	router.DELETE("/delete/:id", controller.Delete())

	err = router.Run()

	if err != nil {
		return
	}

}
