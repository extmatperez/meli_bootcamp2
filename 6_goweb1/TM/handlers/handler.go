package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Code        string    `json:"code"`
	IsPublished bool      `json:"isPublished"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Transaction struct {
	ID              int       `json:"id"`
	TransactionCode string    `json:"transactionCode"`
	Coin            string    `json:"coin"`
	Amount          float64   `json:"amount"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	CreatedAt       time.Time `json:"createdAt"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Height    float64   `json:"height"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

type AllProducts struct {
	Products []Product `json:"products"`
}

type AllTransactions struct {
	Transactions []Transaction
}

type AllUsers struct {
	Users []User
}

// InitHandler responds with a greeting
func InitHandler(ctx *gin.Context) {

	name := ctx.Query("name")

	if name == "" {
		ctx.JSON(http.StatusBadRequest, "Hello user! You must send a name through Query!")
	}

	greet := fmt.Sprintf("Hello %s!", name)

	ctx.JSON(200, greet)
}

func GetAllProducts(ctx *gin.Context) {
	data, err := os.ReadFile("./products.json")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var response AllProducts

	json.Unmarshal(data, &response)

	ctx.JSON(200, gin.H{
		"products": response.Products,
	})
}

func GetAllTransactions(ctx *gin.Context) {
	data, err := os.ReadFile("./transactions.json")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var response AllTransactions

	json.Unmarshal(data, &response)

	ctx.JSON(http.StatusOK, gin.H{
		"transactions": response.Transactions,
	})
}

func GetAllUsers(ctx *gin.Context) {
	data, err := os.ReadFile("./users.json")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	var response AllUsers

	json.Unmarshal(data, &response)

	ctx.JSON(http.StatusOK, gin.H{
		"users": response.Users,
	})
}
