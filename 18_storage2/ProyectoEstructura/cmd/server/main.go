package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/18_storage2/ProyectoEstructura/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/18_storage2/ProyectoEstructura/docs"
	producto "github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/18_storage2/ProyectoEstructura/internal/producto"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/18_storage2/ProyectoEstructura/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//@title Meli Api
//@version 1.0
//@description Api CRUD de productos
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

	db := store.NewStore(store.FileType, "./dbProductos.json")
	repo := producto.NewRepository(db)
	service := producto.NewService(repo)
	controller := handler.NewProductoController(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	groupProducts := router.Group("api/productos")
	{
		groupProducts.GET("/", controller.GetAll())
		groupProducts.POST("/", controller.Store())
		groupProducts.PUT("/:id", controller.Update())
		groupProducts.DELETE("/:id", controller.Detele())
		groupProducts.PATCH("/:id", controller.UpdateNameAndPrice())

	}

	errRun := router.Run()
	if errRun != nil {
		log.Fatal("No se pudo ejectura router.Run()")
	}
}
