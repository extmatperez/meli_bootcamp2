package main

import "fmt"

type Person struct {
	name      string
	employed  bool
	age       uint8
	salary    float32
	seniority float32
}

func main() {
	clients := []Person{
		{
			name:      "Ticiano",
			employed:  true,
			age:       30,
			salary:    100000,
			seniority: 3,
		},
		{
			name:      "Raul",
			employed:  true,
			age:       40,
			salary:    200000,
			seniority: 5,
		},
		{
			name:      "Martina",
			employed:  true,
			age:       25,
			salary:    140000,
			seniority: 2,
		},
		{
			name:      "Nicolas",
			employed:  true,
			age:       21,
			salary:    80000,
			seniority: 1,
		},
		{
			name:      "Maria",
			employed:  false,
			age:       22,
			salary:    0,
			seniority: 0,
		},
	}
	for _, client := range clients {
		if client.age >= 22 &&
			client.seniority >= 1 && client.employed == true {
			if client.salary >= 100000 {
				fmt.Printf("El cliente %v califica para un prestamo con intereses", client.name)
			} else {
				fmt.Printf("El cliente %v califica para un prestamo sin intereses", client.name)
			}

		}
	}
}
