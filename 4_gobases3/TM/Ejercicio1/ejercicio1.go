package main

import (
	"fmt"
	"os"
)

func main() {
	var products []Product

	// Create products with random numbers
	for i := 1; i < 5; i++ {
		id := i
		price := float64(i) * float64(i) * 50.5
		amount := i * 10
		products = append(products, Product{id: id, price: price, amount: amount})
	}

	var textToWrite string
	for i := 0; i < len(products); i++ {
		productData := fmt.Sprintf("%d;%.2f;%d\n", products[i].id, products[i].price, products[i].amount)
		textToWrite += productData
	}

	err := os.WriteFile("../products.txt", []byte(textToWrite), 0644)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Producto escrito correctamente")
	}
}

type Product struct {
	id     int
	price  float64
	amount int
}
