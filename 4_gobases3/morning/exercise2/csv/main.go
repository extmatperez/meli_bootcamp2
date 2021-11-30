package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Producto struct {
	ID       string
	Precio   string
	Cantidad string
}

func main() {
	csvFile, err := os.Open("../../archivo/myFile1.csv")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, csvLine := range csvLines {
			list := Producto{
				ID:       csvLine[0],
				Precio:   csvLine[1],
				Cantidad: csvLine[2],
			}
			fmt.Printf("%-10v%10v%10v\n", list.ID, list.Precio, list.Cantidad)
		}
	}

}
