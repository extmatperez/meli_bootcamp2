/*
Una empresa que se encarga de vender productos de limpieza necesita:
1-Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
separados por punto y coma (csv).
2-Debe tener el id del producto, precio y la cantidad.
3-Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/
package main

import (
	"fmt"
	"os"
)

type Product struct {
	IdProduct int     `json:"id"`
	Price     float64 `json:"precio"`
	Quantity  int     `json:"cantidad"`
}

func saveFileTXT(productList []Product) {
	convCSV := "ID;Price;Quantity\n"
	// productListTransformer = productList
	for _, prod := range productList {
		convCSV += fmt.Sprintf("%v;%v;%v \n", prod.IdProduct, prod.Price, prod.Quantity)
	}
	/*
		// Marshal no reconoce el salto de linea \n por lo que no se puede ocupar en este caso
		productListByte, err2 := json.Marshal(convCSV)
		if err2 != nil {
			fmt.Println("No se pudo escribir")
		}
		fmt.Print(productListByte)
	*/
	err := os.WriteFile("a.csv", []byte(convCSV), 0644)
	if err != nil {
		fmt.Println("No se pudo escribir")
	}

}

func main() {
	producto1 := Product{IdProduct: 1, Price: 90.90, Quantity: 10}
	producto2 := Product{IdProduct: 2, Price: 250.20, Quantity: 4}
	producto3 := Product{IdProduct: 3, Price: 50.31, Quantity: 80}

	var productList []Product

	productList = append(productList, producto1)
	productList = append(productList, producto2)
	productList = append(productList, producto3)

	saveFileTXT(productList)

}
