package main

// Chocolate Company
import (
	"errors"
	"fmt"
)

func main() {
	var name string
	var cant int
	fmt.Println("¿Como se llama el alumno?: ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("¿Cuantas calitifaciones tiene el alumno?: ")
	fmt.Scanf("%d\n", &cant)
	var notes []int
	var note int

	for i := 0; i < cant; i++ {
		fmt.Println("Ve introduciendo las notas: ")
		fmt.Scanf("%d\n", &note)
		notes = append(notes, note)
	}
	fmt.Println(notes)

	result, err := getAverage(name, notes...)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func getAverage(student string, notes ...int) (result string, err error) {
	var sum int
	var cant int
	for _, a := range notes {
		if a < 0 {
			err = errors.New("La nota no puede ser negativo")
			return "", err
		}
		sum += a
		cant++
	}
	average := (float64(sum)) / (float64(cant))
	return "El promedio del alumno " + student + " es " + fmt.Sprintf("%.2f", average), nil
}
