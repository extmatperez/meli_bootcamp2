/*
La misma empresa necesita leer el archivo almacenado, para ello requiere
que: se imprima por pantalla mostrando los valores tabulados,
con un título (tabulado a la izquierda para el ID y a la derecha para el
Precio y Cantidad), el precio, la cantidad y abajo del precio se debe
visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50

*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	IdProduct int     `json:"id"`
	Price     float64 `json:"precio"`
	Quantity  int     `json:"cantidad"`
}

func saveFileTXT(productList []Product) {

	// productListTransformer = productList
	productListByte, err := json.Marshal(productList)
	if err != nil {
		fmt.Println("No se pudo escribir")
	}
	fmt.Print(productListByte)
	err2 := os.WriteFile("a.txt", productListByte, 0644)
	if err2 != nil {
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

	//Read File
	// data, err := os.ReadFile("/Users/joserios/Desktop/bootcamp/meli_bootcamp2/4_gobases3/a.txt")
	// if err == nil {
	// 	file := string(data)
	// 	fmt.Println(file)
	// } else {
	// 	fmt.Println("El archivo no existe")
	// }

}
