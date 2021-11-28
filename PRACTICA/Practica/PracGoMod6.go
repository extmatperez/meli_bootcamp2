package main

import (
	"errors"
	"fmt"
)

func main() {
	var m = map[string]int{}
	// var n = make(map[string]int)
	// var s = map[string]int{"nico": 1}
	// fmt.Println(m, n, s, len(m), len(n), len(s))
	// fmt.Println(s["nico"])

	m["cele"] = 23
	// fmt.Println(m)

	m["matilde"] = 3
	// fmt.Println(m)

	delete(m, "cele")
	fmt.Println(m)

	m["cele"] = 23
	for k, v := range m {
		fmt.Println(k, v)
	}

	suma := sumar(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(suma)

	divi, err := dividir(1.5, 0.0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(divi)
	}

	c := Circulo{radio: 5}
	c.setRadio(5)
}

func sumar(valores ...int) int {
	var r int

	for _, v := range valores {
		r = r + v
	}
	return r
}

func dividir(dividendo, divisor float64) (float64, error) {
	var r float64
	if divisor == 0 {
		return r, errors.New("error: El divisor no puede ser 0")
	} else {
		r = dividendo / divisor
		return r, nil
	}
}

type Circulo struct {
	radio float64
}

func (c *Circulo) setRadio(r float64) {
	c.radio = r
	fmt.Println(r)
}
