package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	list := [][]string{
		{"ID", "Precio", "Cantidad"},
		{"1", "100", "10"},
		{"2", "200", "20"},
		{"3", "300", "30"},
		{"4", "400", "40"},
	}
	csvFile, err := os.Create("../../archivo/myFile1.csv")

	if err != nil {
		fmt.Println(err)
	} else {
		csvWriter := csv.NewWriter(csvFile)
		for _, value := range list {
			csvWriter.Write(value)
		}
		csvWriter.Flush()
		csvFile.Close()
	}
}
