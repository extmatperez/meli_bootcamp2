//Una empresa que se encarga de vender productos de limpieza necesita:
//Implementar una funcionalidad para guardar un archivo de texto, con la
//información de productos comprados, separados por punto y coma (csv).
//Debe tener el id del producto, precio y la cantidad.
//Estos valores pueden ser hardcodeados o escritos en duro en una variable.

package main

import (
	"fmt"
	"os"
)

func main() {

	prod1 := Product{111223, 30012.00, 1}
	prod2 := Product{444321, 1000000.00, 4}
	prod3 := Product{434321, 50.50, 1}

	prodSlice := []Product{prod1, prod2, prod3}

	newCompany := Company{prodSlice}

	/*csvFile, err := os.Create("products.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	productos := [][]string{}
	productos = append(productos, fmt.Sprint("ID; Price; Quantity\n"))
	for _, prod := range newCompany.Products {

		if err != nil {
			fmt.Printf("error")
		} else {
			productos = append(productos, fmt.Sprintf("%d; %f; %d", prod.ID, prod.Price, prod.Quantity))
		}

	}
	csvwriter.Write(productos)
	csvwriter.Flush()
	csvFile.Close() */

	stringToWriteInFile := "ID;Price;Quantity\n"

	for _, prod := range newCompany.Products {
		stringToWriteInFile += fmt.Sprintf("%v;%10.2f;%v\n", prod.ID, prod.Price, prod.Quantity)
	}

	os.WriteFile("./products.csv", []byte(stringToWriteInFile), 0644)

	fmt.Println(newCompany)
}

type Company struct {
	Products []Product
}

type Product struct {
	ID       int
	Price    float64
	Quantity int
}
