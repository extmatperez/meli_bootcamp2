package main

import "fmt"

func main() {

	str := "Hello"
	fmt.Println("Cantidad de Letras ", len(str))

	fmt.Println("Deletreamos: ")
	for _, elem := range str {
		fmt.Printf("%c\n", elem)
	}

}
