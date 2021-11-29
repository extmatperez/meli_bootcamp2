package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {

	minFunc, err1 := operation(minimum)
	valueMinimum := minFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Print("------- Smaller Number --------")
	if err1 != nil {
		fmt.Printf("Error when find min of the numbers: %w", err1)
	} else {
		fmt.Printf("Total smaller value: %v", valueMinimum)
	}

	promFunc, err2 := operation(average)
	valueAverage := promFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Print("------- Average Number --------")
	if err2 != nil {
		fmt.Printf("Error when find avg of the numbers: %w", err2)
	} else {
		fmt.Printf("Total average value: %v", valueAverage)
	}

	maxFunc, err3 := operation(maximum)
	valueMaximum := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Print("------- Smaller Number --------")
	if err3 != nil {
		fmt.Printf("Error when find min of the numbers: %w", err3)
	} else {
		fmt.Printf("Total bigger value: %v", valueMaximum)
	}
}

func operation(value string) (func(values ...int) int, error) {

	switch value {
	case "minimum":
		return minFunc, nil
	case "average":
		return argFunc, nil
	case "maximum":
		return maxFunc, nil
	default:
		return nil, errors.New("There was a error")
	}

}

func minFunc(values ...int) int {

	smallerNumber := values[0]
	for _, number := range values {
		if number < smallerNumber {
			smallerNumber = number
		}
	}
	return smallerNumber

}

func argFunc(values ...int) int {

	var total int
	for _, number := range values {
		total += number
	}
	return total / len(values)
}

func maxFunc(values ...int) int {

	biggerNumber := values[0]
	for _, number := range values {
		if number > biggerNumber {
			biggerNumber = number
		}
	}
	return biggerNumber
}
