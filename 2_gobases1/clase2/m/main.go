package main

import (
	"errors"
	"fmt"
	"math"
)

// Ej1
func calculateTaxes(salary float64) float64 {
	if salary < 50000 {
		return 0
	}
	if salary < 150000 {
		return salary * 0.17
	}
	return salary * 0.27
}

// Ej2
func average(nums ...int) (float64, error) {
	total := 0
	for _, num := range nums {
		total += num
		if num < 0 {
			return -1, errors.New("no acepto negativos")
		}
	}
	return float64(total) / float64(len(nums)), nil
}

// Ej3
func calculateSalary(category string, minutesWorked int) float64 {
	hours := minutesWorked/60 + minutesWorked%60

	if category == "C" {
		return float64(hours * 1000)
	}
	if category == "B" {
		total := hours * 1500
		return float64(total) * 1.2
	}
	total := hours * 3000
	return float64(total) * 1.5
}

// Ej4
func getStatFunc(action string) func(...int) float64 {
	if action == "minimo" {
		return func(nums ...int) float64 {
			min := math.MaxInt
			for _, v := range nums {
				if v < min {
					min = v
				}
			}
			return float64(min)
		}
	}
	if action == "maximo" {
		return func(nums ...int) float64 {
			max := math.MinInt
			for _, v := range nums {
				if v > max {
					max = v
				}
			}
			return float64(max)
		}
	}

	return func(nums ...int) float64 {
		total := 0
		for _, v := range nums {
			total += v
		}
		return float64(total) / float64((len(nums)))
	}
}

func animal(name string) (func(int) float64, error) {
	switch name {
	case "perro":
		return func(quantity int) float64 {
			return float64(quantity * 10)
		}, nil
	case "gato":
		return func(quantity int) float64 {
			return float64(quantity * 5)
		}, nil

	case "hamster":
		return func(quantity int) float64 {
			return float64(quantity) * 0.25
		}, nil

	case "tarantula":
		return func(quantity int) float64 {
			return float64(quantity) * 0.125
		}, nil

	default:
		return func(_ int) float64 {
			return -1.0
		}, errors.New("animal no registrado")
	}
}

func main() {
	taxes := calculateTaxes(200000.0)
	fmt.Println("EJ1")
	fmt.Printf("El impuesto a pagar es de %v\n", taxes)
	fmt.Println("EJ2")
	avg, err := average(1, 2, 3, 4, 5)
	fmt.Println(avg)
	fmt.Println(err)

	avg, err = average(-1, 2, 3, 4, 5)
	fmt.Println(avg)
	fmt.Println(err)

	fmt.Println("EJ3")
	fmt.Println(calculateSalary("C", 203020202))
	fmt.Println("EJ4")

	getAverage := getStatFunc("promedio")

	avg = getAverage(1, 2, 3, 4, 5)
	fmt.Println(avg)

	fmt.Println("EJ5")
	getPerro, _ := animal("perro")

	comidaPerro := getPerro(2)

	fmt.Println(comidaPerro)

	getSerpiente, err := animal("serpiente")
	if err != nil {
		fmt.Println(err)
	} else {
		comidaSerpiente := getSerpiente(2)
		fmt.Println(comidaSerpiente)
	}

}
