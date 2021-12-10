package calculadora

import "fmt"

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("forbidden operation, parameter den must be different of 0")
	}
	return num1 / num2, nil
}
