package main

import (
	"fmt"
	"os"
	"strconv"
)

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		if res, err := strconv.Atoi(v); err == nil {
			numeri = append(numeri,res)
		}
	}
	return
}

func Occorrenze(numeri []int, n int) (occorrenza int) {
	for _, v := range numeri {
		if v == n {
			occorrenza++
		}
	}
	return
}

func main() {
	numeri := LeggiNumeri()
	var somma int

	for _, v := range numeri {
		if Occorrenze(numeri, v) == 1 {
			somma += v
		}
	}
	fmt.Println(somma)
}
