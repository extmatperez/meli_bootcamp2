package main

import (
	"fmt"
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/docs"
	productos "github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/pkg/store"
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
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error al intentar cargar el archivo.env")
	}
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	file := store.New(store.FileType, "./productos.json")
	repo := productos.NewRepository(file)
	service := productos.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pr := r.Group("/productos")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/", p.Edit())
	pr.PATCH("/:id", p.Change())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
