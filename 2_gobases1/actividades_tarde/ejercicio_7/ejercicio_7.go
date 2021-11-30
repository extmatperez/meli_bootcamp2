package main

import "fmt"

var age int
var employee bool
var antiquity int
var salary float32

func main() {
	if age > 22 {
		if employee == true {
			if antiquity > 1 {
				if salary > 100000 {
					fmt.Println("You have a discount ")
				} else {
					fmt.Println("You don`t have the discount based on your salary")
				}
			} else {
				fmt.Println("You don`t have the antiquity yet")
			}
		} else {
			fmt.Println("You don`t are employee")
		}

	} else {
		fmt.Println("You don`t have the age yet")
	}
}
