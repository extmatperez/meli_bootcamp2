package main

import "fmt"

var monthNumber int = 2
var months = map[int]string{
	1: "january", 2: "february", 3: "march", 4: "april", 5: "may", 6: "june",
	7: "july", 8: "august", 9: "september", 10: "october", 11: "november", 12: "december",
}

func main() {
	fmt.Println("Actual month: ", months[monthNumber])
}
