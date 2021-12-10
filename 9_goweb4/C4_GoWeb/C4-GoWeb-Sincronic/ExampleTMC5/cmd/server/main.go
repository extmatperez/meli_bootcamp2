package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTMC5/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTMC5/docs"
	users "github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTMC5/internal/users"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTMC5/pkg/store"
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
	err := godotenv.Load("/Users/joserios/Desktop/bootcamp/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTMC5/cmd/server/.env")
	if err != nil {
		log.Fatal("error al intentar cargar archivo")
	}
	router := gin.Default()

	db := store.New(store.FileType, "./users.json")

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routerUser := router.Group("/users")
	routerUser.GET("/get", controller.GetAll())
	routerUser.POST("/add", controller.Store())
	routerUser.PUT("/:id", controller.Update())
	routerUser.DELETE("/:id", controller.Delete())
	routerUser.PATCH("/lastname/:id", controller.UpdateLastName())
	routerUser.PATCH("/age/:id", controller.UpdateAge())

	router.Run()
}
