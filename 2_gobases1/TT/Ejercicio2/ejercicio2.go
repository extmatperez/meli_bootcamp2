package main

import "fmt"

func applyDiscount(price float64, discount float64) float64 {
	return price * (100 - discount) / 100
}

func main() {
	finalPrice := applyDiscount(30000, 5)
	fmt.Printf("El precio final es: %.2f\n", finalPrice)
}
