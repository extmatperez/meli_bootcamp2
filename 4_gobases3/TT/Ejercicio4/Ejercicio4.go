package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertion_sort(c chan time.Duration, nums []int) {
	ini_insertion := time.Now()
	var aux int
	for i := 1; i < len(nums); i++ {
		aux = nums[i]
		for j := i - 1; j >= 0 && nums[j] > aux; j-- {
			nums[j+1] = nums[j]
			nums[j] = aux
		}
	}
	fin_insertion := time.Now()
	exec_time := fin_insertion.Sub(ini_insertion)
	c <- exec_time
}

func bubble_sort(c chan time.Duration, nums []int) {
	ini_bubble := time.Now()
	var aux int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				aux = nums[i]
				nums[i] = nums[j]
				nums[j] = aux
			}
		}
	}
	fin_bubble := time.Now()
	exec_time := fin_bubble.Sub(ini_bubble)
	c <- exec_time
}

func selection_sort(c chan time.Duration, nums []int) {
	ini_selection := time.Now()
	for i := 0; i < len(nums); i++ {
		min, pos_min := nums[i], i
		original := nums[i]
		for j := i + 1; j < len(nums); j++ {
			aux := nums[j]
			if aux < min {
				min, pos_min = aux, j
			}
		}

		if min != original {
			nums[i], nums[pos_min] = min, original
		}
	}
	fin_selection := time.Now()
	exec_time := fin_selection.Sub(ini_selection)
	c <- exec_time
}

func main() {
	c100_ins := make(chan time.Duration)
	c100_bub := make(chan time.Duration)
	c100_sel := make(chan time.Duration)

	c1000_ins := make(chan time.Duration)
	c1000_bub := make(chan time.Duration)
	c1000_sel := make(chan time.Duration)

	c10000_ins := make(chan time.Duration)
	c10000_bub := make(chan time.Duration)
	c10000_sel := make(chan time.Duration)

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	go insertion_sort(c100_ins, variable1)
	go bubble_sort(c100_bub, variable1)
	go selection_sort(c100_sel, variable1)

	go insertion_sort(c1000_ins, variable2)
	go bubble_sort(c1000_bub, variable2)
	go selection_sort(c1000_sel, variable2)

	go insertion_sort(c10000_ins, variable3)
	go bubble_sort(c10000_bub, variable3)
	go selection_sort(c10000_sel, variable3)

	var exec_c100_ins = <-c100_ins
	var exec_c100_bub = <-c100_bub
	var exec_c100_sel = <-c100_sel

	var exec_c1000_ins = <-c1000_ins
	var exec_c1000_bub = <-c1000_bub
	var exec_c1000_sel = <-c1000_sel

	var exec_c10000_ins = <-c10000_ins
	var exec_c10000_bub = <-c10000_bub
	var exec_c10000_sel = <-c10000_sel

	fmt.Println("\nTiempo de ejecucion arreglo de 100 con ordenamiento por insercion: ", exec_c100_ins)
	fmt.Println("\nTiempo de ejecucion arreglo de 100 con ordenamiento por burbuja: ", exec_c100_bub)
	fmt.Println("\nTiempo de ejecucion arreglo de 100 con ordenamiento por seleccion: ", exec_c100_sel)

	fmt.Println("\nTiempo de ejecucion arreglo de 1000 con ordenamiento por insercion: ", exec_c1000_ins)
	fmt.Println("\nTiempo de ejecucion arreglo de 1000 con ordenamiento por burbuja: ", exec_c1000_bub)
	fmt.Println("\nTiempo de ejecucion arreglo de 1000 con ordenamiento por seleccion: ", exec_c1000_sel)

	fmt.Println("\nTiempo de ejecucion arreglo de 10000 con ordenamiento por insercion: ", exec_c10000_ins)
	fmt.Println("\nTiempo de ejecucion arreglo de 10000 con ordenamiento por burbuja: ", exec_c10000_bub)
	fmt.Println("\nTiempo de ejecucion arreglo de 10000 con ordenamiento por seleccion: ", exec_c10000_sel)
}
