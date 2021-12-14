package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	var data = [][]string{{"1", "2000", "4"}, {"2", "3560", "3"}}

	file, err := os.Create("product.csv")

	if err != nil {
		fmt.Println("No se pudo crear el archivo")
	} else {
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		for _, value := range data {
			err := writer.Write(value)
			if err != nil {
				break
			}
		}
	}

}
