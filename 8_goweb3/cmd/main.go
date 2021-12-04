package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/internal/products"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("no al intentar cargar archivo .env")
	}

	router := gin.Default()

	db := store.NewStore("file", "products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	productsroute := router.Group("/products")

	productsroute.GET("/", controller.GetAll())
	//products.GET("/:id", getProductbyID)
	// products.GET("/products/filter/select", getbyFilter)

	productsroute.POST("/addproduct", controller.AddProduct())
	productsroute.PUT("/updateproduct/:id", controller.UpdateProduct())

	router.Run()

}
