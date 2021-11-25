package main

import (
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

	stock_string := "ID;Price;Amount\n"
	for _, product := range stock {
		stock_string += fmt.Sprintf("%v;%v;%v;\n", product.ID, product.Price, product.Amount)
	}

	err := os.WriteFile("./stock.csv", []byte(stock_string), 0644)
	if err == nil {
		fmt.Println("archivo cargado con exito")
	} else {
		fmt.Println("no pudo cargarse el archivo")
	}
}
