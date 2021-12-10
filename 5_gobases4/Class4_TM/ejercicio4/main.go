package main

import "fmt"

func main() {
	horas_trabajadas := 85
	precio_hora := 1000.0
	salario, err := calc_salario(horas_trabajadas, precio_hora)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salarios es de %.1f con una cantidad de %d horas trabajadas\n", salario, horas_trabajadas)
	}

	map_salarios := map[int]float64{1: 124000.34, 2: 43500, 3: 108900, 4: 136000}
	salario_may, aguinaldo, err := calc_aguinaldo(map_salarios)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("El aguinaldo es de %.1f, y el mejor salario fue de %.1f\n", aguinaldo, salario_may)

}

func calc_salario(horas int, valor float64) (float64, error) {
	sueldo := 0.0
	if horas == 0 {
		return 0, fmt.Errorf("error: No se han registrado horas trabajadas")
	}

	if horas < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		sueldo = float64(horas) * valor
		if sueldo >= 150000 {
			sueldo = sueldo - (sueldo * 0.1)
		}
	}
	return sueldo, nil
}

func calc_aguinaldo(salarios map[int]float64) (float64, float64, error) {
	salario := 0.0
	for _, val := range salarios {
		if val > salario {
			salario = val
		}
	}

	aguinaldo := salario / 12.0 * float64(len(salarios))
	return salario, aguinaldo, nil
}
