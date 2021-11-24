package main

import "fmt"

func main() {

	var word string
	fmt.Println("ingrese la palabra a evaluar")
	fmt.Scanf("%s", &word)
	var long int = len(word)
	fmt.Printf("Ejercicio 1\n")
	fmt.Printf("the word is %v characters long\n", long)

	for i, letter := range word {
		fmt.Printf(" %d index => %s letter \n", i, string(letter))
	}

}
