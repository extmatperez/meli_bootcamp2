package main

func promedio()(notas ...float64) float64 {

	var prom float64

	for _, nota := range notas {
		prom += nota
	}
	prom = prom / float64(len(notas))
	return prom
}
