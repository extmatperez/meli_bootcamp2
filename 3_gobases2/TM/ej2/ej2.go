package ej2

import (
	"errors"
	"fmt"
)

func Ej2(notes ...float64) (float64, error) {
	result := 0.0

	for _, el := range notes {
		if el < 0 {
			fmt.Println(result)
			fmt.Println("Error")
			return result, errors.New("numero negativo")
		} else {
			result += el
		}
	}
	fmt.Println(result)
	return result, nil
}
