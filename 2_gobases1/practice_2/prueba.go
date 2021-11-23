package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a [3]int
	fmt.Printf("a ")
	var b []string
	fmt.Printf("b", b)

	var s = []bool{true, false}
	fmt.Printf("\n%v", s)

	slice := make([]int, 5)
	fmt.Printf("\narray: %v, %T", a, a)
	fmt.Printf("\nslice: %v, %T", slice, slice)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4])

	slice2 := make([]int, 5)

	slice2 = append(slice2, 2)
	fmt.Println(slice2)

	fmt.Printf("\nslice: %d, len: %d, cap: %d", slice2, len(slice), cap(slice))

	slice2 = append(slice2, 3)
	slice2 = append(slice2, 4)
	slice2 = slice2[2 : len(slice)-1]

	fmt.Print("\nsizeof(): %d %T", unsafe.Sizeof(slice2), slice2)

	var myMap = map[string]int{}

	fmt.Printf("\nmyMap %v %T", myMap, myMap)
	fmt.Printf("\nmyMap['Digneli'] %v", myMap["Digneli"])
	myMap["Digneli"] = 27
	myMap["Abel"] = 30
	myMap["Maria"] = 40
	myMap["Abel"] = 50
	fmt.Printf("\nmyMap %v %T", myMap, myMap)
	fmt.Printf("\nmyMap['Digneli'] %v", myMap["Digneli"])

	var students = map[string]int{"Benjamin": 20, "Dig": 30}
	fmt.Printf("\nstudents %v %T", students, students)
	delete(students, "Abel")
	fmt.Printf("\nstudents %v %T", students, students)
	delete(students, "Benjamin")
	fmt.Printf("\nstudents %v %T", students, students)

	students["Maria"] = 40
	students["Abel"] = 50

	array := make([]int, 5)

	array = append(array, 1)
	array = append(array, 2)
	array = append(array, 3)

	for key, value := range students {
		fmt.Printf("\n key %v, value %v", key, value)
	}

}
