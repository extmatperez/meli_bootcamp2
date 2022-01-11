package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/extlurosell/meli_bootcamp_go_w2-2/cmd/server/routes"
	"github.com/extlurosell/meli_bootcamp_go_w2-2/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func tokenCheck() gin.HandlerFunc {

	fmt.Println("PASAMOS POR EL MIDDLEWARE")

	secretToken := os.Getenv("SECRET_TOKEN")

	return func(c *gin.Context) {

		sendToken := c.GetHeader("token")

		if sendToken == "" {
			c.AbortWithStatusJSON(400, gin.H{
				"mensaje": "No se envio token",
			})
			return
		}

		if sendToken != secretToken {
			c.AbortWithStatusJSON(401, gin.H{
				"mensaje": "Token no valido",
			})
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
		fmt.Println("Error en el dotenv", err)
	}
	//db,  := sql.Open("mysql", "meli_sprint_user:MeliSprint#123@/melisprint")
	db, _ := sql.Open("mysql", "test_db_user:Test_DB#123@/testing_db")
	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(tokenCheck())
	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
