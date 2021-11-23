package main

import "fmt"

func main() {
	var age, approvement uint8
	var has_employees bool
	var labor_seniority float32
	var salary float64
	var answer string

	// Client data
	age= 20
	has_employees = true
	labor_seniority = 1.5
	salary = 150000

	if age < 22 {
		approvement += 1
		answer = "You're too young, try again in a few years"
	}

	if !has_employees && labor_seniority < 1 {
		approvement += 1
		answer = "You're not had the labor seniority required for the credit"
	}

	if salary > 100000 && approvement < 1{
		fmt.Println("You are a candidate for an interest-free loan and ")
	}

	if approvement < 1 {
		answer = "Your credit was approved"
	}

	println(answer)

}