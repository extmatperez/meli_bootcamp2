package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("ejecución finalizada")
	}()

	data, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(data)
	}

}
