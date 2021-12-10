package main

import "fmt"

func main() {

	fmt.Println("1 exercise: White box vs Black box")
	differenceBlackWhiteBox()
	fmt.Println("2 exercise: test funcional")
	fmt.Println("Para ejecutar estas pruebas no es necesario el conocimiento de la estructura interna del software")
	fmt.Println("es una especie de caja negra basada en ejecucion, revision y retroalimentacion de las funcionalidades previamente diseñadas. Estas pruebas se hacen mediante el diseño de modelos de prueba que buscan evaluar cada una de las opciones con las que cuenta el programa. Es decir son pruebas especificas, concretas y exhaustivas para probar y validar que el software hace lo que debe y sobre todo, lo que se ha especificado.")
	fmt.Println("3 exercise: test de integracion")
	fmt.Println("Se centra es probar la comunicacion entre componentes y sus comunicaciones ya sea hardware o software")
	fmt.Println("4 exercise: dimensiones de calidad en MELI")
	fmt.Println(". Debe haber mas de 80% de cobertura de calidad de codigo")
	fmt.Println(". Tener el cuenta el tiempo en la toma de decisiones, diseño del testing")
	fmt.Println(". El output del test debe ser realmente relevante")
	fmt.Println(". Tratar de tener actitud de aprendiz")
}

func differenceBlackWhiteBox() {

	blackBoxDefinition()
	whiteBoxDefinition()
	fmt.Println("Black: lo lleva a cabo probadores y White: la realizan los desarrolladores de software")
	fmt.Println("Black: no se requieren conocimientos de implementacion ni programacion y White: si se requieren esos conocimientos")
	fmt.Println("Black: aplicable en niveles mas alto de pruebas como de sistema o aceptacion y White: aplicacion mas baja para pruebas unitaras y de integracion")
	fmt.Println("Black: significa prueba funcional o externa y White: significa prueba estructural o interna")
	fmt.Println("-------Objetivos:--------")
	fmt.Println("Black: las pruebas se concentran principalmente en la funcionalidad del sistema bajo prueba.")
	fmt.Println("White: las pruebas se concentran principalmente en prueba del codigo del programa del sistema bajo prueba, como la estructura del codigo, las ramas, las condiciones, los bucles, etc.")
	fmt.Println("Black: verifica que funcionalidad esta realizando el sistema bajo prueba")
	fmt.Println("White: comprobar el rendimiento del sistema")
}

func blackBoxDefinition() {
	fmt.Println("Black box: Es un metodo de prueba de software sin conocer la estructura interna del codigo o programa")
}

func whiteBoxDefinition() {
	fmt.Println("White box: Es el metodo de prueba en donde el evaluador conoce la estructura interna")
}
