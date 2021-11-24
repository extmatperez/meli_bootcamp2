package main

import "fmt"

func main() {
	var edad int
	var empleado string
	var antiguedad int
	var sueldo int

	fmt.Println("Ejercicio 3")
	fmt.Println("Buenos dias que edad tiene?")
	fmt.Scanf("%d", &edad)
	fmt.Println("Esta usted empleado actualemente?")
	fmt.Scanf("%s", &empleado)

	if empleado == "si" {
		fmt.Println("hace cuantos años esta usted empleado?")
		fmt.Scanf("%d", &antiguedad)
		fmt.Println("cual es su sueldo actualmente?")
		fmt.Scanf("%d", &sueldo)
	}

	if edad >= 22 && empleado == "si" && antiguedad >= 1 {
		if sueldo >= 100000 {
			fmt.Println("felicidades su perfil crediticio califica para poder tener un prestamo sin el cobro de intereses")
		} else {
			fmt.Println("felicidades su perfil crediticio califica para poder tener un prestamo, luego se le comunicara el cobro de intereses")
		}
	} else {
		fmt.Println("lamentablemente su perfil crediticio no califica para tener un prestamo en este banco")
	}

}
