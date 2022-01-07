package handler

import (
	"os"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/web"
	"github.com/gin-gonic/gin"
)

func ValidaToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")

	if token != os.Getenv("TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "Token invalido"))
		return false
	}
	return true
}
