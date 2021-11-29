package main

import (
	"errors"
	"fmt"
)

// Vamos a hacer que nuestro programa sea un poco más complejo y útil.
// Desarrolla las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
// Dicha función deberá retornar más de un valor (salario calculado y error).
// En caso de que el salario mensual sea igual o superior a $150.000, se le
// deberá descontar el 10% en concepto de impuesto.
// En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
// la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede
// haber trabajado menos de 80 hs mensuales”.
// Calcular el medio aguinaldo correspondiente al trabajador
// Fórmula de cálculo de aguinaldo:
// [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
// La función que realice el cálculo deberá retornar más de un valor,
// incluyendo un error en caso de que se ingrese un número negativo.
func calcularSalarioSegunHorasTrabajadas(horasTrabajadas int, valorHora float64) (float64, error) {
	if valorHora <= 0 || horasTrabajadas < 80 {
		return 0, errors.New("ERROR: el trabajador no puede haber trabajado menos de 8 horas")
	}
	subtotal := float64(horasTrabajadas) * valorHora
	if subtotal < 150000 {
		subtotal = subtotal - subtotal/10
	}
	return subtotal, nil

}
func calcularAguinaldo(mejorSalarioSemestre float64, mesesTrabajados int) (float64, error) {
	if mejorSalarioSemestre <= 0 || mesesTrabajados <= 0 {
		return 0, errors.New("ERROR: los parametros no pueden ser numeros negativos")
	}
	return mejorSalarioSemestre / 12 * float64(mesesTrabajados), nil
}

type CustomError struct {
	msg   string
	valor int
}

func (e CustomError) Error() error {
	return fmt.Errorf("%s %d", e.msg, e.valor)
}
func devolver_error(salario int, mensaje string) error {
	var e CustomError
	e.msg = mensaje
	e.valor = salario
	return e.Error()
}

func main() {
	salary := 30
	if salary < 150.0 {
		fmt.Println(devolver_error(salary, "error: el minimo imponible es de 150.000 y el salario ingresado es de :"))
	} else {
		fmt.Println("debe pagar impuesto")
	}
}
