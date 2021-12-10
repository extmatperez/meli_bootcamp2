package calculator

import "errors"

func Subtract(num1 int, num2 int) int {
	return num1 - num2
}

func Divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor can't be 0")
	}

	return dividend / divisor, nil
}
