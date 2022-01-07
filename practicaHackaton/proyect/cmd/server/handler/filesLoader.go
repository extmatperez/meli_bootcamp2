package handler

import (
	loader "github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/internal/filesLoader"
	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/web"
	"github.com/gin-gonic/gin"
)

type FilesLoader struct {
	service loader.Service
}

func NewFilesLoader(ser loader.Service) *FilesLoader {
	return &FilesLoader{service: ser}
}

func (c *FilesLoader) StoreCustomers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !ValidaToken(ctx) {
			return
		}

		err := c.service.CustomersLoader()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, "todo ok", ""))
	}
}
