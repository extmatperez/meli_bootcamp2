package main

import (
	"fmt"
	"errors"
)

// Calcular promedio
func avgNotes(calif ...float64) (float64, error) {
	result := 0.0
	sum := 0.0
	for _, c := range calif {
		sum += c;

		if c < 0 {
			return 0, errors.New("Error");
		};
	}

	result = (sum / float64(len(calif)));
	return result, nil;
}

func main() {
	res, err := avgNotes(5,8,3)
	//res, err := avgNotes(5,8,-3)

	if err != nil {
		fmt.Println("Uno de los números ingresados es negativo")
	} else {
		fmt.Println("Promedio:", res)
	}
}