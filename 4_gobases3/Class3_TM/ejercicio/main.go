package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	Id       = "ID"
	Precio   = "Precio"
	Cantidad = "Cantidad"
)

func main() {
	p1 := producto{1, 342, 20}
	p2 := producto{2, 562, 10}
	p3 := producto{3, 985, 25}
	p4 := producto{4, 345, 32}

	var lista []producto

	lista = append(lista, p1, p2, p3, p4)
	//[]byte(lista)

	productoFormated, err := json.Marshal(lista)

	err = os.WriteFile("./archivo.txt", productoFormated, 0644)

	data, err := os.ReadFile("./archivo.txt")
	var pListaLeida []producto
	json.Unmarshal(data, &pListaLeida)

	if err != nil {
		fmt.Printf("\nSe encontro el archivo: \n")
	} else {
		// file := string(data)
		//fmt.Println()
		// fmt.Println(file)
		fmt.Printf("\n%s %18s %10s\n", Id, Precio, Cantidad)
		for i := range pListaLeida {
			// persona := string(p)
			// fmt.Printf("\n %T\n", p)
			fmt.Printf("%-10d %10.1f %10d\n", pListaLeida[i].Id, pListaLeida[i].Precio, pListaLeida[i].Cantidad)
		}
	}

	// fmt.Printf("\n%+v", pListaLeida)

}

type producto struct {
	Id       int
	Precio   float64
	Cantidad int
}
