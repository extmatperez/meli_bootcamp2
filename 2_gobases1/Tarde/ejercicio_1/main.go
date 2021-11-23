package main

import (
	"fmt"
	"strings"
)

func main() {
	palabra := "como vas"
	var letras []string
	letras = strings.Split(palabra, "")
	fmt.Printf("Número de letras: %v \n", len(letras))
	for _, letra := range letras {
		fmt.Println(letra)
	}

}
