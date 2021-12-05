package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/projectC4/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/projectC4/docs"
	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/projectC4/internal/products"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/projectC4/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
		log.Fatal("error al intentar cargar archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	router.PUT("/products/:id", controller.Update())
	// router.PATCH("/products/:id", controller.UpdateProd())
	// router.DELETE("/products/:id", controller.Delete())
	router.Run()

}
