package main

import (
	"fmt"
	"os"
	"strconv"
)

func Calcola(sl []int) (somma int) {
	for i,v1 := range sl {
		for _,v2:= range sl[i+1:] {
			if(v1*v2) % 2 == 0{
				somma+= v1*v2
			}
		}
	}
	return somma
}

func main() {
	var toCalculate []int

	toCalculate = make([]int, len(os.Args[1:]))

	for _, arg := range os.Args[1:] {

		if str, err := strconv.Atoi(arg); err == nil {
			toCalculate = append(toCalculate, str)
		}
	}
	fmt.Println("La somma è:", Calcola(toCalculate))
}
