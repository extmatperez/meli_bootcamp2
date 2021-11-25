package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	list := [][]string{
		{"ID", "PRECIO", "CANT"},
		{"1", "157.99", "20"},
		{"2", "25.33", "100"},
		{"3", "115", "500"},
	}
	csvFile, err := os.Create("../Files/ProdList.csv")
	if err == nil {
		csvwriter := csv.NewWriter(csvFile)
		for _, empRow := range list {
			csvwriter.Write(empRow)
		}
		csvwriter.Flush()
		csvFile.Close()
	} else {
		fmt.Print("Error al abrir archivo")
	}
}
