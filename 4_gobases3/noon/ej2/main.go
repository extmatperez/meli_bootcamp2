package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Id    int     `json: "id"`
	Price float64 `json: "price"`
	Count int     `json: "count"`
}

func main() {
	productsJson, err := os.ReadFile("../ej1/products.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var products []Product
	json.Unmarshal(productsJson, &products)

	fmt.Printf("%-10s \t\t %10s \t\t %10s\n\n", "ID PRODUCT", "PRICE", "QUANTITY")
	var total float64 = 0.0
	for _, product := range products {
		total = total + product.Price
		fmt.Printf("%-10d \t\t %10.2f \t\t %10d\n", product.Id, product.Price, product.Count)
	}
	fmt.Printf("%10s \t\t %10.2f \t\t %10s\n", "", total, "")
	fmt.Println()

	var productsBytes []byte
	productsBytes, err = os.ReadFile("../ej1/productsParsed.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%-10s \t\t %10s \t\t %10s\n\n", "ID PRODUCT", "PRICE", "QUANTITY")
	total = 0
	productsParsed := strings.Split(string(productsBytes), ";")
	for i := 0; i < len(productsParsed)-1; i = i + 3 {
		fmt.Printf("%-10s \t\t %10s \t\t %10s", productsParsed[i], productsParsed[i+1], productsParsed[i+2])
	}
	fmt.Println()
}
