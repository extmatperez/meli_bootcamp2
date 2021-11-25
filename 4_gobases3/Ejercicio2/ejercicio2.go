package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Producto struct {
	id       int
	precio   float64
	cantidad int
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}
func main() {
	records := readCsvFile("./archivos/datos.csv")
	for j, reg := range records {
		if j == 0 {
			fmt.Printf("%-10v%-10v%-10v\n", reg[0], reg[1], reg[2])
		} else {
			i, _ := strconv.Atoi(reg[0])
			p, _ := strconv.ParseFloat(reg[1], 64)
			aux := reg[2]
			aux2 := aux[:len(aux)-1]
			c, _ := strconv.Atoi(aux2)
			list := Producto{
				id:       i,
				precio:   p,
				cantidad: c,
			}
			fmt.Printf("%-10v%-10.2f%-10v\n", list.id, list.precio, list.cantidad)
		}
	}
}
