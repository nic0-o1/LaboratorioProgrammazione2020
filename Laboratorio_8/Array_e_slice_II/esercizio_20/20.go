package main

import (
	"fmt"
	"os"
	"strconv"
)

func Calcola(sl []int) (res int) {
	fmt.Println(sl)
	for i := 0; i < len(sl); i += 2 {
		res += sl[i] * sl[i+1]
	}
	return
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
