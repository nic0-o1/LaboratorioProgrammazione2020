package main

import (
	"sort"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numero := os.Args[1]

	primi := GeneraPrimi(numero)

	sort.Ints(primi)

	for i := range primi {
		fmt.Println(primi[i])
	}
}

func ÈPrimo(n int) bool {
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GeneraPrimi(numero string) (primi []int) {

	n, _ := strconv.Atoi(numero)

	if ÈPrimo(n) {
		fmt.Println(n)
	}

	for lunghezza := 1; lunghezza <= 3; lunghezza++ {
		for posizione := range numero[:len(numero)-lunghezza+1] {
			sottostringa := numero[:posizione] + numero[posizione+lunghezza:]
			n, _ := strconv.Atoi(sottostringa)
			if ÈPrimo(n) {
				fmt.Println(n)
				primi = append(primi, n)
			}
		}
	}
	fmt.Println()
	return
}


