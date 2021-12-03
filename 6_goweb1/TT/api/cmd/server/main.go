package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/docs"
	producto "github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/internal/producto"
	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	db := store.New(store.FileType, "../../internal/producto/products.json")

	repository := producto.NewRepository(db)
	service := producto.NewService(repository)
	controller := handler.NewProduct(service)

	routerGroup := router.Group("/products")
	{
		//routerGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		routerGroup.GET("", controller.GetAll())
		routerGroup.POST("/add", controller.Store())
		routerGroup.DELETE("/:id", controller.Delete())
		routerGroup.PUT("/:id", controller.Update())
		routerGroup.PATCH("/:id", controller.UpdateNombre())
	}

	router.Run()
}
