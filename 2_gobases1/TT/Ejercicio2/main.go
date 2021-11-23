package main

import "fmt"

func main() {
	var price float64 = 100000
	discount := 0.15

	final_price := (price - (price * discount))

	fmt.Println("The discount is: ", (discount * 100),"%")
	fmt.Println("The value of the discount is: $", (discount * price))
	fmt.Println("The final price is: $", final_price)
}