package main

import (
	"fmt"
	"log"
	"os"

	handler "github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/docs"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/customers"
	loader "github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/filesLoader"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/invoicers"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/products"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/sales"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")
	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("api_token")
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}
		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
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
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	repoCustomers := customers.NewRepository()
	serviceCustomers := customers.NewService(repoCustomers)
	customersHandler := handler.NewCustomer(serviceCustomers)

	repoInvocers := invoicers.NewRepository()
	serviceInvoicers := invoicers.NewService(repoInvocers)
	invoicersHandler := handler.NewInvoicer(serviceInvoicers)

	repoProducts := products.NewRepository()
	serviceProducts := products.NewService(repoProducts)
	productsHandler := handler.NewProduct(serviceProducts)

	repoSales := sales.NewRepository()
	serviceSales := sales.NewService(repoSales)
	salesHandler := handler.NewSale(serviceSales)

	repoLoader := loader.NewRepository()
	serviceLoader := loader.NewService(repoLoader)
	filesLoaderHandler := handler.NewFilesLoader(serviceLoader)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(TokenAuthMiddleware())

	customers := router.Group("/customers")
	{
		customers.GET("/", customersHandler.GetAll())
		customers.GET("/:id", customersHandler.GetByID())
		customers.POST("/", customersHandler.Store())
		customers.PUT("/:id", customersHandler.Update())
		customers.DELETE("/:id", customersHandler.Delete())
	}

	invoicers := router.Group("/invoicers")
	{
		invoicers.GET("/", invoicersHandler.GetAll())
		invoicers.GET("/:id", invoicersHandler.GetByID())
		invoicers.POST("/", invoicersHandler.Store())
		invoicers.PUT("/:id", invoicersHandler.Update())
		invoicers.DELETE("/:id", invoicersHandler.Delete())
	}

	products := router.Group("/products")
	{
		products.GET("/", productsHandler.GetAll())
		products.GET("/:id", productsHandler.GetByID())
		products.POST("/", productsHandler.Store())
		products.PUT("/:id", productsHandler.Update())
		products.DELETE("/:id", productsHandler.Delete())
	}

	sales := router.Group("/sales")
	{
		sales.GET("/", salesHandler.GetAll())
		sales.GET("/:id", salesHandler.GetByID())
		sales.POST("/", salesHandler.Store())
		sales.PUT("/:id", salesHandler.Update())
		sales.DELETE("/:id", salesHandler.Delete())
	}

	filesLoader := router.Group("/filesLoader")
	{
		filesLoader.POST("/customers", filesLoaderHandler.StoreCustomers())
		filesLoader.POST("/invoicers", filesLoaderHandler.StoreInvoicers())
		filesLoader.POST("/products", filesLoaderHandler.StoreProducts())
		filesLoader.POST("/sales", filesLoaderHandler.StoreSales())
	}

	err = router.Run(":8080")

	if err != nil {
		fmt.Println(err)
	}
}
