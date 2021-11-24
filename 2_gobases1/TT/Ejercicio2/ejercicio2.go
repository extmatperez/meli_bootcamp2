package main

import "fmt"

var impo int = 2500
var desc int = 25

func main() {
	importFinal := (impo * desc) / 100

	impo = impo - importFinal

	fmt.Println(impo)
}
