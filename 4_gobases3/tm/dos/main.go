package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID     int
	Price  float64
	Amount int
}

var stock []Product

func main() {

	stock_file, err1 := os.ReadFile("../uno/store_stock.txt")
	if err1 == nil {
		err2 := json.Unmarshal(stock_file, &stock)
		if err2 == nil {
			fmt.Println("archivo leido con exito")
			fmt.Printf("%-10s %10s %s\n", "ID", "Precio", "Cantidad")
			for _, product := range stock {
				fmt.Printf("%-10v", product.ID)
				fmt.Printf("%10v", product.Price)
				fmt.Printf("%10v\n", product.Amount)
			}
		} else {
			fmt.Println("no puso leerse el archivo")
		}
	} else {
		fmt.Println("no pudo recuperarse el archivo")
	}
}
