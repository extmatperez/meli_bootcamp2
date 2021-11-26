/* Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

package main

import "fmt"

type Product struct {
	Product_id int     `json:"nombre"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

func save_product() {

}

func main() {
	fmt.Println("Hello World!")
}

/* package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func guardarArchivo(prod producto, file string) {

	data, err := os.ReadFile(file)

	var ListaProds []producto

	json.Unmarshal(data, &ListaProds)

	ListaProds = append(ListaProds, prod)

	formProdList, err := json.Marshal(ListaProds)

	erro := os.WriteFile(file, formProdList, 0644)

	if err != nil {
		fmt.Println("E
	if err != nil {
		fmt.Println("Error json: ", err)
	}

	if erro != nil {
		fmt.Println("Error", err)
	}
}

func leerArchivo(file string) {
	data, err := os.ReadFile(file)

	var ListaProds []producto

	json.Unmarshal(data, &ListaProds)

	if err == nil {
		formatoPrint(ListaProds)
	} else {
		fmt.Println("El archivo no existe...")
	}
}

func formatoPrint(lista []producto) {
	var total float64
	fmt.Printf("ID\t Precio Cantidad")
	for i := 0; i < len(lista); i++ {
		total += lista[i].Precio * float
Arevalo Ivan13:42
for i := 0; i < len(lista); i++ {
		total += lista[i].Precio * float64(lista[i].Cantidad)
		fmt.Printf("\n%v\t %v\t %v", lista[i].Id, lista[i].Precio, lista[i].Cantidad)
	}
	fmt.Printf("\n\t %v\n", total)

}

func main() {
	// prod1 := producto{1112, 100.50, 3}
	// guardarArchivo(prod1, "../../a.txt")
	leerArchivo("../../a.txt")
} */
