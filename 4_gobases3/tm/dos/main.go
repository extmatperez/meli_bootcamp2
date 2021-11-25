package main

import (
	"fmt"
	"os"
	"strings"
)

type Product struct {
	ID     int
	Price  float64
	Amount int
}

func main() {

	stock_file, err := os.ReadFile("../uno/stock.csv")
	stock := string(stock_file)
	if err == nil {
		fmt.Println("archivo leido con exito")
		newStr := strings.Replace(string(stock), ";", "\t", -1)
		fmt.Println(newStr)
		//for _, product := range stock {
		//	fmt.Printf("%-10v", product.ID)
		//	fmt.Printf("%10v", product.Price)
		//	fmt.Printf("%10v\n", product.Amount)
		//}
	} else {
		fmt.Println("no puso leerse el archivo")
	}
}
