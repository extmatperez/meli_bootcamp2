package middleware

import (
	"os"

	"github.com/extmatperez/meli_bootcamp2/9_goweb4/pkg/web"
	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") == "" {
			ctx.AbortWithStatusJSON(400, web.NewResponse(400, nil, "Missing token"))
			return
		}

		if ctx.GetHeader("token") != os.Getenv("TOKEN") {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Don´t have permission to access"))
			return
		}

		ctx.Next()
	}
}
