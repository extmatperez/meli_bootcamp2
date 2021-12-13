package divide

import "fmt"

func Divide(numer, denom int) (int, error) {
	if denom == 0 {
		return 0, fmt.Errorf("Denominator cannot be zero")
	}
	return numer / denom, nil
}
