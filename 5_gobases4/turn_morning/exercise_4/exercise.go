package main

import "fmt"

func main() {

	sal1 := Salary{-200000.00, 1, 2021, 1}
	sal2 := Salary{250000.00, 2, 2021, 1}
	sal3 := Salary{260000.00, 3, 2021, 1}
	sal4 := Salary{270000.00, 4, 2021, 1}
	sals := []Salary{sal1, sal2, sal3, sal4}
	worker := Worker{sals, 1000, 500000.00}

	total, err := totalSalaryByMonth(worker.hoursWorked, worker.actualSalary)

	if err != nil {
		errSaved := fmt.Errorf("Error numer 1: %v", err)
		fmt.Println(errSaved)
	} else {
		fmt.Println("The total salary by hours was: ", total)
	}

	err2, total2 := getBonusByWorker(1, &worker)

	if err2 != nil {
		errSaved := fmt.Errorf("Error number 2: %v", err2)
		fmt.Println(errSaved)
	} else {
		fmt.Println("Bonus: ", total2)
	}
}

type WorkerError struct {
	Message string `json:"message"`
	Code    int    `miEtiqueta:"code"`
}

func (err *WorkerError) Error() string {
	return "Interface error: " + err.Message
}

type Worker struct {
	Salaries     []Salary `json:"salaries"`
	hoursWorked  int
	actualSalary float64
}

type Salary struct {
	total    float64 `json:"cant"`
	month    int     `json:"month"`
	year     int     `json:"year"`
	semester int     `json:"semester"`
}

func totalSalaryByMonth(hours int, salary float64) (float64, error) {
	total := float64(hours) * salary

	if hours < 80 {
		return 0.0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	if total >= 150000 {
		total += .90
	}
	return total, nil

}

func getBonusByWorker(semester int, worker *Worker) (error, float64) {
	var bestSalary float64
	var monthsWorker int
	for _, v := range worker.Salaries {
		if v.total < 0 {
			return &WorkerError{"error: ", 400}, 0.00
		}
		if semester == 1 {
			if bestSalary < v.total {
				bestSalary = v.total
			}
			monthsWorker = v.month
		} else if semester == 2 {
			if bestSalary < v.total {
				bestSalary = v.total
			}
			switch v.month {
			case 5:
				monthsWorker = 1
			case 6:
				monthsWorker = 2
			case 7:
				monthsWorker = 3
			case 8:
				monthsWorker = 4
			}
		}
	}

	result := bestSalary / 12 * float64(monthsWorker)

	return nil, result
}
