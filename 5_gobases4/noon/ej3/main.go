package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const baseSalary int = 150000

func calcTax(salary int) (string, error) {
	if salary < baseSalary {
		return "", fmt.Errorf("the salary is less than %d", baseSalary)
	}

	return "You have to pay taxes", nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	salary := rand.Intn(100000) + 100000

	msg, err := calcTax(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(msg)
}
