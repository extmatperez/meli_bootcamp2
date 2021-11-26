package main

import (
	"fmt"
	"os"
	"strings"
)

type producto struct {
	Id       string `json:"id"`
	Precio   string `json:"precio"`
	Cantidad string `json:"cantidad"`
}

func read() {
	var prodSalida []producto
	data, _ := os.ReadFile("./myfile.txt")
	s := strings.Split(string(data), "\n")

	for i := 0; i < len(s); i++ {
		s2 := strings.Split(s[i], ";")

		pr1 := &producto{Id: s2[0], Precio: s2[1], Cantidad: s2[2]}

		prodSalida = append(prodSalida, *pr1)
	}

	fmt.Println("data  ", prodSalida)
	fmt.Printf("\n %-15v", "ID")
	fmt.Printf(" %15v", "Precio")
	fmt.Printf(" %10v", "Cantidad")
	for i := 0; i < len(prodSalida); i++ {
		fmt.Printf("\n %-15v", prodSalida[i].Id)
		fmt.Printf(" %15v", prodSalida[i].Precio)
		fmt.Printf(" %10v", prodSalida[i].Cantidad)
		fmt.Println("")
	}

}

func main() {
	read()
}
