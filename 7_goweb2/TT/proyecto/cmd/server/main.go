// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/docs"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/products"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	db := store.New(store.FileType, "./productsList.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)
	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	router.PUT("/products/:id", controller.Update())
	router.PATCH("/products/:id", controller.UpdateProd())
	router.DELETE("/products/:id", controller.Delete())
	router.Run()
}
