package main

import "fmt"

var price float64
var discount float64
var result float64

func main() {
	price = 2000
	discount = 0.10
	result = price * discount
	fmt.Printf("The price is: $%v \n", price)
	fmt.Printf("The discount is: %v%% \n", discount*100)
	fmt.Printf("The discount value is: $%v \n", result)
	fmt.Printf("The total result is: $%v \n", price-result)
}
