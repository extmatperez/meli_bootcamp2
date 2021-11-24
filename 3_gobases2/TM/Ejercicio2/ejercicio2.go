package main

import (
	"errors"
	"fmt"
)

func main() {
	average, err := getAverage(10, 9, 5, 4, 5, 6, 7, 4, 2, 6, 8, 9, 2, 1, 10, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio es: %.2f\n", average)
	}
}

func getAverage(scores ...int) (float64, error) {
	var sum int = 0
	for i := 0; i < len(scores); i++ {
		if scores[i] < 0 {
			return 0.0, errors.New("no puede haber notas negativas")
		}
		sum += scores[i]
	}

	return float64(sum) / float64(len(scores)), nil
}
