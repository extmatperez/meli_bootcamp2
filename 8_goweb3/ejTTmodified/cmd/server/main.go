package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/8_goweb3/ejTTmodified/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/8_goweb3/ejTTmodified/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/8_goweb3/ejTTmodified/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/ncostamagna/meli-bootcamp/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title PRODUCTS API
// version 1.0
// @description API wich can handle products with a CRUD
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldnt load the environment")
	}
	router := gin.Default()
	storage := store.NewStore("file", "products.json")
	repository := internal.NewRepository(storage)
	service := internal.NewService(repository)
	producto := handler.NewProducto(service)
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/producto", producto.GetAll())
	router.POST("/producto", producto.Store())
	router.GET("/producto/:id", producto.GetProductById())
	router.PUT("/producto/:id", producto.Update())
	router.DELETE("/producto/:id", producto.Delete())
	router.PATCH("/producto/:id", producto.UpdateNombrePrecio())

	router.Run()

}
