package main

import (
	"fmt"
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTMM/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTMM/docs"
	product "github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTMM/internal/products"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTMM/pkg/store"
	"github.com/extmatperez/meli_bootcamp2/tree/castillo_walter/8_goweb3/go-web-TTMM/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title WALTER PRODUCTS API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load("./8_goweb3/go-web-TTMM/cmd/server/.env")
	if err != nil {
		fmt.Println(err)
		log.Fatal("error al intentar cargar el archivo")
	}

	router := gin.Default()
	db := store.New(store.FileType, "./8_goweb3/go-web-TTMM/cmd/server/products.json")
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	controller := handler.NewProduct(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/product/get", TokenAuthMiddleware(), controller.GetAll())
	router.POST("/product/add", controller.Store())
	router.PUT("/product/put/:id", controller.Update())
	router.PATCH("/product/patch/:id", controller.UpdateNombre())
	router.PATCH("/product/patchprecio/:id", controller.UpdatePrecio())
	router.DELETE("/product/delete/:id", controller.Delete())
	router.Run()
}

//middleware
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no hay variable de entorno de token cofigurada")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			respondWithError(ctx, 400, web.NewReponse(400, nil, "falta token"))
			//ctx.JSON(400, web.NewReponse(400, nil, "falta token"))
			return
		}
		tokenEnv := os.Getenv("TOKEN")
		if token != tokenEnv {
			respondWithError(ctx, 400, web.NewReponse(400, nil, "token incorrecto"))
			//ctx.JSON(400, web.NewReponse(404, nil, "token incorrecto"))
			return
		}
		ctx.Next()
	}
}
