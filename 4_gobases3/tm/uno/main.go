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
	stock = append(stock, Product{1, 24.99, 3})
	stock = append(stock, Product{2, 8.99, 7})
	stock = append(stock, Product{3, 15.50, 4})

	stock_file, err1 := json.Marshal(stock)

	if err1 == nil {
		err2 := os.WriteFile("./store_stock.txt", stock_file, 0644)
		if err2 == nil {
			fmt.Println("archivo cargado con exito")
		} else {
			fmt.Println("no pudo cargarse el archivo")
		}
	} else {
		fmt.Println("la conversion fallo")
	}

}
