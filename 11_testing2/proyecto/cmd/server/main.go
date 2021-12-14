package main

import (
	"log"
	"net/http"
	"os"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/docs"
	internal "github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/internal/transactions"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//Middleware
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			RespondWithError(ctx, http.StatusUnauthorized, "No se ingresó un token")
			return
		}
		if token != requiredToken {
			RespondWithError(ctx, http.StatusUnauthorized, "Token incorrecto")
			return
		}
		ctx.Next()
	}
}

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handles MELI Transactions.
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

	db := store.New(store.FileType, "./transactions.json")
	repo := internal.NewRepository(db)
	service := internal.NewService(repo)
	controller := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "pong")
	// })
	router.Use(TokenAuthMiddleware())
	routerTransactions := router.Group("/transactions") //, controller.ValidateToken())
	{
		routerTransactions.GET("/", controller.GetAll())
		routerTransactions.GET("/:id", controller.GetTransactionByID())
		routerTransactions.POST("/", controller.Store())
		routerTransactions.PUT("/:id", controller.Update())
		routerTransactions.DELETE("/:id", controller.Delete())
		routerTransactions.PATCH("/:id", controller.UpdateCodigoYMonto())
	}
	err = router.Run()
	if err != nil {
		log.Fatal("Error en el servidor.")
	}
	// //List Endpoints
	// rutas := router.Routes()
	// for _, r := range rutas {
	// 	fmt.Printf("Method: %s \t\t Path: %s\n", r.Method, r.Path)
	// }

}
