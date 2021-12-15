package test

import (
	"os"

	"github.com/extmatperez/meli_bootcamp2/12_testing3/turn_afternoon/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/meli-bootcamp/cmd/server/handler"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./products.json")
	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	return r
}

//Update

//Delete
