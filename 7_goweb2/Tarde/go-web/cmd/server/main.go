package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/docs"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/internal/products"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.htm
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el ambiente")
	}
	db := store.New("file", os.Getenv("FILEPATH"))
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	controller := handler.NewProduct(service)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productsRoute := router.Group("products")
	productsRoute.GET("", controller.GetAll())
	productsRoute.GET("/filter", controller.Filter())
	productsRoute.GET("/:id", controller.FindById())
	productsRoute.POST("", controller.Store())
	productsRoute.PUT("/:id", controller.Update())
	productsRoute.DELETE("/:id", controller.Delete())
	productsRoute.PATCH("/:id", controller.UpdateNameAndPrice())

	router.Run()
}
