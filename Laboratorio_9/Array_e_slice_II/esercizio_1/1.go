package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func MinimoDispari(sl []int) int {
	min := math.MaxInt64
	for _, v := range sl {
		if v%2 != 0 && v < min {
			min = v
		}
	}
	return min
}

func main() {
	var numeri []int
	for _, in := range os.Args[1:] {
		if val, err := strconv.Atoi(in); err == nil {
			numeri = append(numeri, val)
		}
	}

	min := MinimoDispari(numeri)

	for _, v := range numeri {
		if v%2 == 0 && v > min {
			fmt.Printf("%v ",v)
		}
	}
}
