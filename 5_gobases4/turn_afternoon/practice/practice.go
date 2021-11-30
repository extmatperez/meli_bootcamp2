package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	prube(a)

	fmt.Println("init..")

	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(err)

	//animals := []string{"vaca", "perro", "halcon"}
	//fmt.Println("solo vuela el ", animals[len(animals)])

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*5)

	<-ctx.Done()
	defer fmt.Println(ctx.Err().Error())
}

func prube(array [5]int) {

	array[4] = 45
}
