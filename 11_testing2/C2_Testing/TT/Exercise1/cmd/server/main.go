package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/11_testing2/C2_Testing/TT/Exercise1/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/11_testing2/C2_Testing/TT/Exercise1/internal/users"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/C2_Testing/TT/Exercise1/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/Users/joserios/Desktop/bootcamp/meli_bootcamp2/11_testing2/C2_Testing/TT/Exercise1/cmd/server/.env")
	if err != nil {
		log.Fatal("error al intentar cargar archivo")
	}
	router := gin.Default()

	db := store.New(store.FileType, "./users.json")

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	routerUser := router.Group("/users")
	routerUser.GET("/get", controller.GetAll())
	routerUser.POST("/add", controller.Store())
	routerUser.PUT("/:id", controller.Update())
	routerUser.DELETE("/:id", controller.Delete())
	routerUser.PATCH("/lastname/:id", controller.UpdateLastName())
	routerUser.PATCH("/age/:id", controller.UpdateAge())

	router.Run()
}
