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

	fmt.Printf("%20.1f", Total(pListaLeida))
}

type producto struct {
	Id       int
	Precio   float64
	Cantidad int
}

func Total(pListaLeida []producto) float64 {
	total := 0.0
	for _, producto := range pListaLeida {
		total += producto.calc_total()
	}
	return total
}

func (p *producto) calc_total() float64 {
	tot_prod := p.Precio * float64(p.Cantidad)
	return tot_prod
}
