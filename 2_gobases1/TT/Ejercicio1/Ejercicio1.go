package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	string_buffer := bufio.NewReader(os.Stdin)
	var word string
	fmt.Println("Palabra: ")
	word, _ = string_buffer.ReadString('\n')
	var letter_array []string = strings.Split(word, "")
	fmt.Printf("La palabra tiene %d letras.", len(letter_array)-1)
	fmt.Println("\nPalabra deletreada: ")
	for _, letter := range letter_array {
		fmt.Printf("\n%s", letter)
	}
}
