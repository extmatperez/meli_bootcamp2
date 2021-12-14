package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("el archivo indicado no fue encontrado o está dañado")
	fmt.Println("init..")

	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}

}
