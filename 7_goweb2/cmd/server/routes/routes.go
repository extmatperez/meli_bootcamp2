/*
	Author: AG-Meli - Andrés Ghione
*/

package routes

import (
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/cmd/server/handler"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/internal/transactions"
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/pkg/store"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db store.Store
}

func NewRouter(r *gin.Engine, db store.Store) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildTransactionsRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildTransactionsRoutes() {
	repo := transactions.NewRepository(r.db)
	service := transactions.NewService(repo)
	handler := handler.NewTransaction(service)
	routerTransactions := r.rg.Group("/transactions")
	{
		routerTransactions.GET("/GetAll", handler.GetAll())
		routerTransactions.GET("/GetByID/:ID", handler.GetByID())
		routerTransactions.POST("/", handler.Store())
		routerTransactions.PUT("/:ID", handler.Update())
		routerTransactions.DELETE("/:ID", handler.Delete())
		routerTransactions.PATCH("/transactionCode/:ID/:TransactionCode", handler.ModifyTransactionCode())
		routerTransactions.PATCH("/amount/:ID/:Amount", handler.ModifyAmount())
	}
}
