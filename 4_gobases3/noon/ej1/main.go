package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Product struct {
	Id    int     `json: "id"`
	Price float64 `json: "price"`
	Count int     `json: "count"`
}

func NewProduct() Product {
	var Id int = rand.Intn(100)
	var Price float64 = (rand.Float64() * 1000) + 1
	var Count int = rand.Intn(500) + 1

	return Product{Id, Price, Count}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var textParsed string

	var products []Product
	product := NewProduct()
	for product.Id != 0 {
		products = append(products, product)
		productText := fmt.Sprintf("%d; %.2f; %d;\n", product.Id, product.Price, product.Count)
		textParsed = fmt.Sprintf("%v%v", textParsed, productText)
		product = NewProduct()
	}

	text, err := json.MarshalIndent(products, "", "\t")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	err = os.WriteFile("./products.txt", text, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile("./productsParsed.txt", []byte(textParsed), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
