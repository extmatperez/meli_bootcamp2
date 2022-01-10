package calculadora

import "fmt"

func Dividir(num1, num2 int) int {
	res := 0
	if num2 == 0 {
		fmt.Println("no se puede dividir por 0")
	} else {
		res = num1 / num2
	}
	return res
}
