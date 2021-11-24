package main

import (
	"fmt"
	"strings"
)

// Needed to use split

var word string = "hola Mundo"

func main() {
	fmt.Printf("The word has %v letters. \n", len(word))
	split := strings.Split(word, "")
	fmt.Println(split)
}
