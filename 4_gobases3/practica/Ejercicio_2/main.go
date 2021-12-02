package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
Id 			int 	`json:"id"`
Price 		int `json:"price"`
Quantity 	int 	`json:"quantity"`
}

func main() {

	data, err := os.ReadFile("../productBD.txt")

	if err != nil {
		fmt.Println("No se pudo leer")
	}

	var pListRead []Product
	json.Unmarshal(data, &pListRead)

	fmt.Println(pListRead)
	fmt.Printf("ID")
	fmt.Printf("%15v %15v\n", "PRICE", "QUANTITY")
	for i, _ := range pListRead {
		fmt.Printf("%v %15v %14v\n", pListRead[i].Id, pListRead[i].Price, pListRead[i].Quantity)
	}
}