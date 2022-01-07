/*
	Author: AG-Meli - Andrés Ghione
*/

package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/cmd/server/routes"
	stores "github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	db := stores.New(stores.FileType, "7_goweb2/internal/transactions/transactions.json")
	r := gin.Default()
	router := routes.NewRouter(r, db)
	router.MapRoutes()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
