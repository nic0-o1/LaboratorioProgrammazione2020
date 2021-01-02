package main

import (
	"fmt"
	"os"
	"strconv"
)

func DivisoriPropri(n int) (divisori []int) {
	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			divisori = append(divisori, i)
		}
	}
	return
}

func DivisoriComuni(divisoriA, divisoriB []int) (inComune int) {
	for _, v := range divisoriA {
		for _, k := range divisoriB {
			if k == v {
				inComune++
			}
		}
	}

	return
}

func ÈPerfetto(n int) bool {
	var somma int
	for _, v := range DivisoriPropri(n) {
		somma += v
	}
	return somma == n
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	divisoriMin, _ := strconv.Atoi(os.Args[2])
	divisoriMax, _ := strconv.Atoi(os.Args[3])
	
	for a := 1; a <= n; a++ {
		for b := a; b <= n; b++ {
			divCom := DivisoriComuni(DivisoriPropri(a), DivisoriPropri(b))
			if (ÈPerfetto(a) || ÈPerfetto(b)) && divCom >= divisoriMin && divCom <= divisoriMax {
				fmt.Println(a, b)
			}
		}
	}

}
