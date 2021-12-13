package main

import (
	"errors"
	"fmt"
)

func main() {
	var operacion int
	fmt.Println("Ingrese la operacion que desea realizar: \n\t1- Calcular sueldo\n\t2- Calcular aguinaldo\n\t3- Salir")
	fmt.Scanf("%d", &operacion)
	switch {
	case operacion == 1:
		show_salary_calculator()
	case operacion == 2:
		show_aguinaldo_calculator()
	case operacion == 3:
		return
	default:
		fmt.Println("No ha seleccionado una opcion valida")
	}
	fmt.Println("Gracias por usar nuestros servicios!")
}

func show_salary_calculator() {
	var horas_trabajadas float64
	var valor_hora float64
	fmt.Println("Ingrese las horas trabajadas por el empleado: ")
	fmt.Scanf("%f", &horas_trabajadas)
	fmt.Println("Ingrese el valor hora del empleado:")
	fmt.Scanf("%f", &valor_hora)
	salary, err := salary_calculator(horas_trabajadas, valor_hora)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	} else {
		fmt.Printf("Su salario es: %.2f\n", salary)
	}
}

func show_aguinaldo_calculator() {
	var best_salary float64 = 0.0
	var months_worked float64 = 0.0

	fmt.Println("\nCalculemos su aguinaldo")
	fmt.Println("Ingrese el mejor sueldo de los ultimos 12 meses: ")
	fmt.Scanf("%f", &best_salary)
	fmt.Println("Ingrese cuantos de esos 12 meses trabajo: ")
	fmt.Scanf("%f", &months_worked)

	aguinaldo, err := aguinaldo_calculator(best_salary, months_worked)
	if err != nil {
		fmt.Printf("%v\n", errors.Unwrap(err))
		return
	} else {
		fmt.Printf("Su aguinaldo es: %.2f\n", aguinaldo)
	}
}

func salary_calculator(horas_trabajadas, valor_hora float64) (float64, error) {
	salario := horas_trabajadas * valor_hora
	switch {
	case salario > 150000:
		salario = salario - (salario * 0.1)
	case horas_trabajadas < 80:
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	return salario, nil
}

func aguinaldo_calculator(best_salary, months_worked float64) (float64, error) {
	if best_salary < 0 || months_worked < 0 {
		err := &ErrorPersonalizado{}
		err2 := fmt.Errorf("%w", err)
		return 0, err2
	} else {
		aguinaldo := (best_salary / 12) * months_worked
		return aguinaldo, nil
	}
}

type ErrorPersonalizado struct{}

func (e ErrorPersonalizado) Error() string {
	return "error: se ha ingresado un numero negativo"
}
