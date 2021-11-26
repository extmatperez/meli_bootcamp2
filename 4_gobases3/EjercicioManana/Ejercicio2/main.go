package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leerArchivo(path string) {

	datos, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Error archivo no existe")
	} else {
		totalPrecio := 0.0
		contenido := string(datos)
		filas := strings.Split(contenido, "\n")
		for i := 0; i < len(filas); i++ {
			columnas := strings.Split(filas[i], ";")
			if i == 0 {
				fmt.Printf("%-10s %10s %10s", columnas[0], columnas[1], columnas[2])
			} else {
				if len(columnas) == 3 {
					precio, err := strconv.ParseFloat(columnas[1], 64)
					if err != nil {
						precio = 0.0
					}

					cantidad, err := strconv.Atoi(columnas[2])

					if err != nil {
						cantidad = 0
					}

					totalPrecio += precio * float64(cantidad)

					fmt.Printf("%-10s %10.2f %10d", columnas[0], precio, cantidad)
				}

			}
			if i < len(filas)-1 {
				fmt.Printf("\n")
			}

		}
		fmt.Printf(" %20.2f\n", totalPrecio)
	}

}

func main() {

	leerArchivo("../Ejercicio1/prueba1.csv")
}
