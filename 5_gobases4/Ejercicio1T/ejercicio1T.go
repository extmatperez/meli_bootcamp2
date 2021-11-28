package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./archivo/customers.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("termino la ejecucion")
}
