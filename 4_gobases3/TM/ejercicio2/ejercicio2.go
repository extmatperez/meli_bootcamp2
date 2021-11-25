package main

/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)
*/

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
		total := 0.00
		for _, line := range csvLines {
			aux, err := strconv.ParseFloat(line[1], 32)
			if err != nil {
				fmt.Print("Error con el numero")
			}
			total += aux
			list := Producto{
				ID:     line[0],
				PRECIO: line[1],
				CANT:   line[2],
			}
			fmt.Printf("%-10v%10v%10v\n", list.ID, list.PRECIO, list.CANT)
		}
		fmt.Printf("%-10v%10.2f%10v\n", " ", total, "")
	} else {
		fmt.Println(err)
	}

}
