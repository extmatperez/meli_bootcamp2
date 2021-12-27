package main

import (
	"log"
	"os"

	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API Documentation
// @version 1.0
// @description This API Handle MELI USers

// @contact.name Archuby Federico
// @contact.email federico.archuby@mercadolibre.com
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't open .env file")
	}

	db := store.NewStore(store.FileType, "./users.json")

	router := gin.Default()
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	hand := handler.NewUser(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := router.Group("/users")
	group.POST("", hand.Store())
	group.GET("", hand.GetAll())
	group.PUT("/:id", hand.Update())
	group.PATCH("/:id", hand.UpdateLastNameAge())
	group.DELETE("/:id", hand.Delete())

	err = router.Run()
	if err != nil {
		log.Fatal("Can't start server")
	}
}
