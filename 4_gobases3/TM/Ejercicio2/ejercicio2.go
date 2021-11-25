package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("../products.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(bytes), "\n")
	var totalPrice float64

	fmt.Printf("%-10s %10s %10s\n", "ID", "Precio", "Cantidad")
	for _, line := range lines {
		if line == "" {
			continue
		}

		values := strings.Split(line, ";")
		price, errConvert := strconv.ParseFloat(values[1], 64)

		if errConvert != nil {
			continue
		}

		totalPrice += price
		fmt.Printf("%-10s %10.2f %10s\n", values[0], price, values[2])
	}

	fmt.Printf("%-10s %10.2f %10s\n", "", totalPrice, "")
}
