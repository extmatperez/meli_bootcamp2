package main

import (
	"fmt"
	"os"
)

//Guardar archivo
type Company struct {
	Products []Product
}

type Product struct {
	ID       int
	Price    float64
	Quantity int
}

func main() {

	prod1 := Product{111223, 30012.00, 1}
	prod2 := Product{444321, 1000000.00, 4}
	prod3 := Product{434321, 50.50, 1}

	prodSlice := []Product{prod1, prod2, prod3}

	newCompany := Company{prodSlice}

	HeadersInFile := "ID;Price;Quantity\n"

	for _, prod := range newCompany.Products {
		HeadersInFile += fmt.Sprintf("%v;%10.2f;%v\n", prod.ID, prod.Price, prod.Quantity)
	}

	os.WriteFile("./prods_comprados.csv", []byte(HeadersInFile), 0644)

	fmt.Println(newCompany)
}