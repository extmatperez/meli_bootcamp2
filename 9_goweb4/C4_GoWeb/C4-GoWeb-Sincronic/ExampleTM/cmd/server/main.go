package main

import (
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTM/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTM/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error al intentar cargar archivo .env")
	// }
	router := gin.Default()

	repo := users.NewRepository()
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
