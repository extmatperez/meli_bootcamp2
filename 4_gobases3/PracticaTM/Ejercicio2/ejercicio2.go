package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../Ejercicio1/ProductosComprados.csv")

	if err != nil {
		fmt.Println(err)
	} else {
		//imprimo el titulo
		fmt.Printf("%-20v%10v%12v\n", "ID", "Precio", "Cantidad")

		//Separo los productos del archivo en la variable
		productosLeidos := strings.Split(string(data), "\n")

		var idProd string
		var precioProd, cantProd, precioFinal float64
		precioFinal = 0
		for i, producto := range productosLeidos {
			//Recorro cada producto y lo imprimo
			productoActual := strings.Split(string(producto), ";")

			if len(productoActual) == 3 && i > 0 {
				idProd = productoActual[0]
				precioProd, _ = strconv.ParseFloat(productoActual[1], 32)
				cantProd, _ = strconv.ParseFloat(productoActual[2], 32)
				fmt.Printf("%-20s%10.2f%12.0f\n", idProd, precioProd, cantProd)

				precioFinal += precioProd * cantProd
			}
		}
		fmt.Printf("%30.2f\n", precioFinal)
	}

}
