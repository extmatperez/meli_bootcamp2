package  main

import (
	"fmt"
)

func main(){
	var (
		name = "Juan"
		age = 23
		work = true
		salary float64 = 195000
	)

	if age < 22 {
		fmt.Println("Lo lamentamos!",name, " No podemos asignar un prestamo a alguien menor a 22 años")
	} else {
		if work == false {
			fmt.Println("Lo lamentamos!",name, "No podemos asignar un credito a alguien desempleado")
		} else {
			if salary > 100000 {
				fmt.Println("Felicidades!",name, "Su credito ha sido aprovado, además queda exento de intereses")
			} else {
				fmt.Println("Felicidades!",name, "su credito ha sido aprovado!")
			}
		}
	}

}