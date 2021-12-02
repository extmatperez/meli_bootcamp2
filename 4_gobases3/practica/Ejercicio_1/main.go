package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id 			int
	Price		int
	Quantity	int
}
func main() {

	p1 := Product{68, 6000, 4}
	p2 := Product{67, 4000, 2}

	var list []Product
	list = append(list, p1)
	list = append(list, p2)

	if err != nil {
		fmt.Println("No se pudo escribir")
	}

}
