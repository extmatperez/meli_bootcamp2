package main

import (
	"errors"
	"fmt"
)

func main() {

	marksAverage, err := avg(7, 8, 9)

	if err != nil {
		fmt.Println("Error encontrado")
	} else {
		fmt.Println(marksAverage)
	}
}

func avg(marks ...int) (float64, error) {

	currentSum := 0

	for _, num := range marks {

		if num < 0 {
			return 0.0, errors.New("nota negativa encontrada")
		}

		currentSum += num
	}

	return float64(currentSum) / float64(len(marks)), nil
}
