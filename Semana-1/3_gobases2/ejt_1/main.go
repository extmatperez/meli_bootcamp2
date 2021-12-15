package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
	Cant  float64 `json:"cant"`
}

func main() {

	var prodList []Product

	p1 := Product{1, 57.30, 100}
	p2 := Product{2, 104.40, 70}
	p3 := Product{3, 43.80, 30}

	prodList = append(prodList, p1)
	prodList = append(prodList, p2)
	prodList = append(prodList, p3)

	codeList, err := json.Marshal(prodList)

	if err == nil {
		err := os.WriteFile("../ejt_1/prodList.txt", codeList, 0644)
		if err == nil {
			fmt.Print("List Charged\n")
		} else {
			fmt.Print("Error")
		}
	} else {
		fmt.Print("List Not Found")
	}
}
