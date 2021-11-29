package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string //
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

// Punto 1. Uso de custom errors.
func (e *myError) devolver_custom_error() error {
	e.msg = "El salario ingresado no alcanza el minimo imponible."
	return e
}

// Punto 2. Uso de errors.New.
func devolver_error() error {
	return errors.New("El salario ingresado no alcanza el minimo imponible.")
}

// Punto 3. Uso de fmt.Errorf.
func devolver_error_errorf() error {
	return fmt.Errorf("El salario ingresado no alcanza el minimo imponible.")
}

// Todas estas son funciones a implementar para el punto 4.
func devolver_salario_mes(hs_work int, hs_price float64) (float64, error) {
	var salary float64 = 0.0
	if hs_work < 80 {
		return 0, fmt.Errorf("El trabajador no puede haber trabajado menos de 80hs mensuales.")
	} else {
		salary = (float64)(hs_work) * hs_price
		if salary > 150000 {
			salary = salary * 0.9
		}
	}
	return salary, nil
}

func devolver_medio_aguinaldo(max_salary float64, months_work int) (float64, error) {
	var aguinaldo float64 = 0.0
	if max_salary <= 0 || months_work <= 0 || months_work > 6 {
		return 0, fmt.Errorf("Valor/es ingresado/s invalido/s.")
	} else {
		aguinaldo = (max_salary / 12) * (float64)(months_work)
	}
	return aguinaldo, nil
}

func main() {
	var salary_calc, hs_price, max_salary, aguinaldo float64
	var salary, hs_work, months_work int
	var err1 myError

	fmt.Println("Ingresa el salario:")
	fmt.Scanf("%d", &salary)
	if salary < 150000 {
		// Punto 1.
		err1_1 := err1.devolver_custom_error()
		// Punto 2.
		err1_2 := devolver_error()
		// Punto 3.
		err1_3 := devolver_error_errorf()
		fmt.Println(err1_1)
		fmt.Println(errors.Unwrap(err1_1))
		fmt.Println(errors.Unwrap(err1_2))
		fmt.Println(errors.Unwrap(err1_3))
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
	fmt.Println("\n================================================================")
	fmt.Println("Ingrese cantidad de horas trabajadas: ")
	fmt.Scanf("%d", &hs_work)
	fmt.Println("Ingrese el precio de las horas trabajadas: ")
	fmt.Scanf("%f", &hs_price)
	salary_calc, err2 := devolver_salario_mes(hs_work, hs_price)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("El salario a liquidar es de $%.2f", salary_calc)
	}
	fmt.Println("\n================================================================")
	fmt.Println("Ingrese salario maximo del semestre: ")
	fmt.Scanf("%f", &max_salary)
	fmt.Println("Ingrese la cantidad de meses trabajados en el semestre: ")
	fmt.Scanf("%d", &months_work)
	aguinaldo, err3 := devolver_medio_aguinaldo(max_salary, months_work)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Printf("El medio aguinaldo a liquidar es de $%.2f", aguinaldo)
	}
}
