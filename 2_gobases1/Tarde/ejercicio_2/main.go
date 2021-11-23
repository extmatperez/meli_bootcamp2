package main

import (
	"fmt"
)

func main() {
	descuento := 50.00
	precio := 50500.00
	total := descuento * precio / 100
	fmt.Printf("Prfecio = %v  Descuento = %v%% \n", precio, descuento)
	fmt.Printf("Valor con descuento = %v \n", total)

}
