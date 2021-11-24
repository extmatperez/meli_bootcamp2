package main

import (
	"fmt"
	"math"
)

func main() {

	var price float64
	var percentage float64
	fmt.Println("Please enter your price: ")
	fmt.Scanf("%f.2\n", &price)
	fmt.Println("Please enter your percentage: ")
	fmt.Scanf("%f.2\n", &percentage)

	var percentageTransformed float64 = (1 - (percentage / 100))
	//fmt.Printf("%f , %T", transforPorcent)
	result := price * percentageTransformed
	percentageRounded := math.Round(percentage)
	fmt.Printf("You have %v %% off and this is the result of your discount: %.2f", percentageRounded, result)

}
