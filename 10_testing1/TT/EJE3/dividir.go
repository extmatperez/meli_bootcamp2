package dividir

import "fmt"

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("No se puede dividir por cero")
	}
	return num1 / num2, nil
}
