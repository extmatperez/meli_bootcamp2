package ej3

import "fmt"

func Ej3(category string, minutes int) float64 {
	salaries := map[string]float64{
		"C": 1000.0 / 60,
		"B": (1500 / 60) * 1.2,
		"A": (3000 / 60) * 1.5,
	}

	fmt.Printf("%.2f\n", float64(minutes)*salaries[category])
	return float64(minutes) * salaries[category]
}
