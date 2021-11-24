package main

import "fmt"

// Descuento
func main() {

	price := 150.0
	discount := 25.0
	final_price := price * (1 - (discount/100))

	fmt.Printf("Final price is: %v\n", final_price)

}
