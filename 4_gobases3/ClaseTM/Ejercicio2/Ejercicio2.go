package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	leerYRenderizar()
}

func leerYRenderizar() {

	data, err := os.ReadFile("../Ejercicio1/ProductosCSVHardcoded.txt")

	if err == nil {
		fmt.Println(data)
		marshFile := ""
		json.Unmarshal(data, &marshFile)

		fmt.Printf("%-10s %-10s %-10s \n", "ID", "PRECIO", "CANTIDAD")

		for i := 0; i < len(marshFile); i++ {
			if string(marshFile[i]) == "/" {
				fmt.Println("")
			} else {

				if string(marshFile[i]) == ";" {
					fmt.Printf("%-5s", "")
				} else {
					fmt.Printf("%s", string(marshFile[i]))
				}
			}
		}
		fmt.Println("")

	} else {
		fmt.Println(err)
	}

}
