package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-TM/Exercise1/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-TM/Exercise1/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {

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
