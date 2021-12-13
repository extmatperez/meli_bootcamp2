package calculadora

import "fmt"

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("el denominador no puede ser cero")
	}
	return num1 / num2, nil
}
