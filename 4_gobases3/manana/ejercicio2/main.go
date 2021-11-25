package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type producto struct {
	Id       string `json:"id"`
	Precio   string `json:"precio"`
	Cantidad string `json:"cantidad"`
}

func read() {
	var prodSalida []producto
	data, _ := os.ReadFile("./myfile.txt")
	json.Unmarshal([]byte(data), &prodSalida)
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
