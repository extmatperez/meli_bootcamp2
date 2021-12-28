package main

import (
	"log"

	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/17_storage1/TT/ejercicio_1/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/tree/parra_diego/17_storage1/TT/ejercicio_1/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/17_storage1/TT/ejercicio_1/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/../docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// meli_bootcamp2/8_goweb3/ejercicio_1_swagger/docs
	db := store.New(store.FileType, "./productoSalida.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	//
	router.PUT("/products/:id", controller.Update())
	router.DELETE("/products/:id", controller.Delete())
	router.PATCH("/products/:id", controller.UpdateNamePrice())

	err2 := router.Run()
	if err2 != nil {
		return
	}
}
