package  main

import "fmt"

func main(){
	var (
		precio float64 = 10000
		desc float64 = 15
	)

	fmt.Println("El precio del producto es: ", precio,
			"\nEl descuento es: ", desc,
			"\nY el precio final es: ", precio-(precio * (desc/100)))

}