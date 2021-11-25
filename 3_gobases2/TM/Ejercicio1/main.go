package main

import "fmt"

func impuesto() {
	var sueldo int = 30000

	switch {
	case sueldo > 50000:
		fmt.Println("17%")
	case sueldo > 150000:
		fmt.Println("10%")
	}
}
