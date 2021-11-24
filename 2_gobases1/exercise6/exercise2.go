package main

import (
	"fmt"
)

// Key : Value -> price : discount
var word_map = map[int]int{1500: 50}

func main() {
	for key, element := range word_map {
		fmt.Println("Precio: ", key, ", El descuento: ", element, "\n el precio final es: ", (key * element / 100))
	}
}
