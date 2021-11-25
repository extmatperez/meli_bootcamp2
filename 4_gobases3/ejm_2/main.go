package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Producto struct {
	ID     string
	PRECIO string
	CANT   string
}

func main() {
	csvFile, err := os.Open("../Files/ProdList.csv")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err == nil {
		for _, line := range csvLines {
			list := Producto{
				ID:     line[0],
				PRECIO: line[1],
				CANT:   line[2],
			}
			fmt.Printf("%-10v%10v%10v\n", list.ID, list.PRECIO, list.CANT)
		}
	} else {
		fmt.Println(err)
	}

}
