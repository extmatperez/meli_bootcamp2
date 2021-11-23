package ej3

import "fmt"

func Ej3(age, yearsWorking, salary int, isEmployeed bool) string {

	answer := ""

	if age > 22 && isEmployeed && yearsWorking >= 1 {
		if salary > 100000 {
			answer = "You can have a loan with no interests!"
		} else {
			answer = "You can have a loan, but you'll have to pay interests!"
		}
	} else {
		answer = "You can't ask a loan yet!"
	}

	fmt.Println(answer)

	return answer
}
