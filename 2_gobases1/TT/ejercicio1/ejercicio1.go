package main

import "fmt"

func main() {
	var name string = "Ticiano Mensegué"

	fmt.Printf("Cantidad de letras: %v", len(name))
	for _, letra := range name {
		fmt.Printf("\n %c", letra)
	}
}
