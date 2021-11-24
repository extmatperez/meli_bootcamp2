package main

import "fmt"

func main() {
	var age int = 23
	var working bool = true
	var antiquity int = 2
	var salary int = 50000

	if age > 22 && working && antiquity > 1 {
		fmt.Print("Puede acceder al prestamo")
		if salary > 100000 {
			fmt.Print(" y ademas sin interes")
		}
	} else {
		fmt.Print("No puede acceder al prestamo")
	}
}
