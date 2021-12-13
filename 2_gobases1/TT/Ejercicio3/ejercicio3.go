package main

import "fmt"

func main() {
	cliente1 := Cliente{"Diego", "Torena", 23, true, 3, 70000}
	for _, m := range dar_prestamos(&cliente1) {
		fmt.Println(m)
	}
}

type Cliente struct {
	Nombre            string `json:"nombre"`
	Apellido          string `json:"apellido"`
	Edad              int    `json:"edad"`
	IsEmployed        bool   `json:"isemployed"`
	AntiguedadLaboral int    `json:"antiguedadlaboral"`
	Sueldo            int    `json:"sueldo"`
}

func dar_prestamos(cliente *Cliente) []string {
	var return_message []string
	if cliente.Edad < 23 {
		return_message = append(return_message, "Debe tener mas de 22 para solicitar un prestamo")
	}
	if cliente.AntiguedadLaboral < 1 {
		return_message = append(return_message, "Debe tener al menos un año de antiguedad para solicitar un prestamo")
	}
	if !cliente.IsEmployed {
		return_message = append(return_message, "Debe tener empleo para solicitar un prestamo")
	}
	if cliente.Sueldo > 100000 && len(return_message) == 0 {
		return_message = append(return_message, "Puede solicitar un prestamo libre de intereses")
	}
	if len(return_message) == 0 {
		return_message = append(return_message, "Puede solicitar su prestamo, pero debera pagar intereses")
	}
	return return_message
}
