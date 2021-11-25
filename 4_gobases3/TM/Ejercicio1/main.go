package main

import (
	"fmt"
	"os"
)

type Producto struct {
	ID       int
	Precio   float64
	Cantidad int
}

type Company struct {
	Product []Producto
}

func main() {
	p1 := Producto{111223, 456.0, 24}
	p2 := Producto{111224, 234.0, 34}
	p3 := Producto{111225, 865.0, 17}
	//p1 := Producto{111226, 456.0, 24}

	prodSlice := []Producto{p1, p2, p3}

	newCompany := Company{prodSlice}

	// fmt.Println(newCompany.)

	stringToWriteInFile := "ID;Price;Cantidad\n"

	for _, prod := range newCompany.Product {
		stringToWriteInFile += fmt.Sprintf("%v;%v;%v\n", prod.ID, prod.Precio, prod.Cantidad)
	}

	os.WriteFile("./newArchive.csv", []byte(stringToWriteInFile), 0644)

}
