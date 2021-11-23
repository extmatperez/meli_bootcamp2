package main

import "fmt"

func main() {
	word := "Prueba"
	length := len(word)

	fmt.Println("The word ", word," has ",length , " letters of length.")
	
	for i := 0; i < len(word); i++ {
		fmt.Println("The value of the letter in the position ", i , " is ", string(word[i]))
	}
}