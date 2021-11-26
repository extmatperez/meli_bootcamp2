package main

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
	fmt.Printf("ID\t  Precio \tCantidad")
	for i := 0; i < len(lista); i++ {
		total += lista[i].Precio * float64(lista[i].Cantidad)
		fmt.Printf("\n%v\t %10.2f\t \t%v", lista[i].Id, lista[i].Precio, lista[i].Cantidad)
	}
	fmt.Printf("\n\t %10.2f\n", total)
}
func main() {
	leerArchivo("./myFile.txt")
}
