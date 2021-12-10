package testing

import "errors"

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("se ingresó un 0 en el denominador")
	}

	return (num / den), nil
}
