// El RECOVER busca un panic, y DEFER hace que se ejecute la funcion al ultimo.
// El RECOVER se ejecuta si o si y no permite que salte en terminal el error proveniente del panic, si bien se corta la ejecucion del programa.
// Hace que no se aborte del todo el proceso.

// CONTEXT.
// Establece un contexto que puede ser pasado por codigo, para que las partes interesadas reaccionen a el.
// Por ejemplo, definir un deadlock, osea, el tiempo maximo de ejecucion de una funcion y si no se termina se corta.
// Si se desea pasar a una funcion, se pasa como el primer argumento, y con ctx.
// Existe por que a veces queremos llamar a una funcion que puede tardar mucho en retornar. Puede pasarse a distintas funciones y a sus "funciones hijas".

// BACKGROUND
// Nos define un contexto vacio.

package main

import (
	"context"
	"fmt"
	"time"
)

func isPair(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if num%2 != 0 {
		panic("No es un numero par!")
	}

	fmt.Println("El numero es par")
}

func saludoWrapper(ctx context.Context) {
	saludo(ctx)
}

func saludo(ctx context.Context) {
	fmt.Println(ctx.Value("Saludo"))
}

func main() {
	// isPair(5)
	// fmt.Println("Termino la ejecucion.")
	ctx := context.Background()
	deadline := time.Now().Add(time.Second * 5)

	// con estas proximas dos definiciones definimos tiempos de corte de funciones hasta que mueren,
	// con deadline (proceso mas largo porque hay que definir deadline antes y timeout, que es definible directo desde la funcion)
	ctx, _ = context.WithDeadline(ctx, deadline)

	ctx, _ = context.WithTimeout(ctx, time.Second*5)

	<-ctx.Done()

	fmt.Println(ctx.Err().Error())

	fmt.Println(ctx)
	ctx = context.WithValue(ctx, "Saludo", "Hola es un saludo en context")
	saludoWrapper(ctx)
}
