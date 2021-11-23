package main

import "fmt"

type Client struct {
	Age        int
	Years      int
	Salary     float64
	IsEmployee bool
}

var clients []Client

func createClient(age, years int, salary float64, isEmployee bool) Client {
	newClient := Client{
		age,
		years,
		salary,
		isEmployee,
	}
	return newClient
}

func main() {
	clients = append(clients, createClient(19, 2, 100000, true))
	clients = append(clients, createClient(27, 5, 189000, true))
	clients = append(clients, createClient(25, 7, 190000, false))
	for _, client := range clients {
		if client.Age >= 22 && client.IsEmployee && client.Years > 1 {
			if client.Salary > 100000 {
				fmt.Println("Client applies for an interest-free loan.")
				continue
			}
			fmt.Println("Client applies for a loan.")
		}
		fmt.Println("Client does not apply for a loan.")
	}
}
