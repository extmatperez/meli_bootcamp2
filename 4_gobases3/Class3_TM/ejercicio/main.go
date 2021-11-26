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
	p5 := producto{0, 345, 0}

	var lista []producto

	lista = append(lista, p1, p2, p3, p4, p5)

	productoFormated, err := json.Marshal(lista)

	err = os.WriteFile("./archivo.txt", productoFormated, 0644)

	data, err := os.ReadFile("./archivo.txt")
	var pListaLeida []producto
	json.Unmarshal(data, &pListaLeida)

	if err != nil {
		fmt.Printf("\nSe encontro el archivo: \n")
	} else {
		fmt.Printf("\n%s %18s %10s\n", Id, Precio, Cantidad)
		for i := range pListaLeida {
			fmt.Printf("%-10d %10.1f %10d\n", pListaLeida[i].Id, pListaLeida[i].Precio, pListaLeida[i].Cantidad)
		}
	}

}

type producto struct {
	Id       int
	Precio   float64
	Cantidad int
}

func (p *producto) calc_total(precio float64, cantidad int) {
	//for
	//return p.Precio * p.Cantidad
}
