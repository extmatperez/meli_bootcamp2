package main

import (
	"fmt"
)

func main() {
	var palabra string = "programar"

	fmt.Println("La palabra '",palabra,"' tiene ",len(palabra)," letras")

	for _, letra := range palabra {
		fmt.Printf("%c\n", letra)
	}
}
