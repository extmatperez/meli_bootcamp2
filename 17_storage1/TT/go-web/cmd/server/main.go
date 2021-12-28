package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/docs"
	products "github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/internal/products"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/TT/go-web/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Falta Token"))
		} else if token != os.Getenv("TOKEN") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "No tiene permisos para realizar la petición solicitada"))

		}

		c.Next()
	}
}

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

	repositorySQL := products.NewRepositorySQL()
	serviceSQL := products.NewServiceSQL(repositorySQL)
	controllerSQL := handler.NewProductSQL(serviceSQL)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(TokenAuthMiddleware())
	productsRoute := router.Group("products")
	productsRoute.GET("", controller.GetAll())
	productsRoute.GET("/filter", controller.Filter())
	productsRoute.GET("/:id", controller.FindById())
	productsRoute.POST("", controllerSQL.Store())
	productsRoute.PUT("/:id", controller.Update())
	productsRoute.DELETE("/:id", controller.Delete())
	productsRoute.PATCH("/:id", controller.UpdateNameAndPrice())

	err = router.Run()
	if err != nil {
		fmt.Println("Error subiendo las rutas")
	}

}
