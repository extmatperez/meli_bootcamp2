package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	useNew        = "new"
	useFmt        = "fmt"
	useUnwrap     = "unwrap"
	minHours      = 80
	errorSalary   = "salary"
	error13Salary = "13salary"
)

type errorHours struct {
	hours int
}

func (e errorHours) Error() string {
	return fmt.Sprintf("The employee can't work less than %d hours", e.hours)
}

func generateError(msg string, otherError error, errorType string) error {
	switch errorType {
	case useNew:
		return errors.New(msg)
	case useFmt:
		return fmt.Errorf(msg)
	case useUnwrap:
		return fmt.Errorf("Error: %w", otherError)
	default:
		return nil
	}
}

func printError(errorType string, msg error) {
	if errorType == useUnwrap {
		fmt.Println(errors.Unwrap(msg))
	} else {
		fmt.Println(msg)
	}
}

func getErrorMessage(errorName string, errorType string) (string, error) {
	var msg string

	switch errorName {
	case errorSalary:
		msg = fmt.Sprintf("The employee can't work less than %d hours", minHours)
	case error13Salary:
		msg = fmt.Sprintf("The employee can't have a negative salary")
	}

	if errorType == useUnwrap {
		return "", errorHours{
			hours: minHours,
		}
	}

	return msg, nil
}

func calculateSalary(hours int, valueHour float64, errorType string) (float64, error) {
	if hours < minHours {
		msg, otherError := getErrorMessage(errorSalary, errorType)
		return 0.0, generateError(msg, otherError, errorType)
	}

	return (float64(hours) * valueHour) * 0.9, nil
}

func get13Salary(salaries []float64, errorType string) (float64, error) {

	var max float64 = 0.0
	for _, salary := range salaries {
		if salary < 0 {
			msg, otherError := getErrorMessage(error13Salary, errorType)
			return 0.0, generateError(msg, otherError, errorType)
		}

		if max < salary {
			max = salary
		}
	}

	var monthsWorked int = rand.Intn(6) + 1

	return (max / 12) * float64(monthsWorked), nil
}

func test13SalaryError(errorType string) {
	var salaries []float64

	for i := 0; i < 12; i++ {
		hoursWorked := rand.Intn(50) + 90
		salary, _ := calculateSalary(hoursWorked, 10.0, errorType)

		var forNegative float64 = rand.Float64()
		if forNegative < 0.05 {
			salary = salary * -1
		}

		salaries = append(salaries, salary)
	}

	newSalary, err := get13Salary(salaries, errorType)
	if err != nil {
		printError(errorType, err)
	} else {
		fmt.Printf("The 13 salary is %.2f\n", newSalary)
	}
}

func testSalaryError(errorType string) {
	hoursWorked := rand.Intn(50) + 70
	salary, err := calculateSalary(hoursWorked, 10.0, errorType)

	if err != nil {
		printError(errorType, err)
	} else {
		fmt.Printf("The salary is %.2f\n", salary)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testSalaryError(useFmt)
	testSalaryError(useNew)
	testSalaryError(useUnwrap)

	fmt.Println("")

	test13SalaryError(useFmt)
	test13SalaryError(useNew)
	test13SalaryError(useUnwrap)
}
