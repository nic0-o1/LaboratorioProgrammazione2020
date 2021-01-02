package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func LeggiNumero() int {
	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		return n
	}
	return 0
}

func CreaSliceBidimensionale(n int) [][]int {
	var s [][]int
	s = make([][]int, n)
	for i := 0; i < n; i++ {
		s[i] = make([]int, n)
	}
	return s
}

func InizializzaSliceBidimensionale(s [][]int) {

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}

}

func Coppie(s [][]int) (coppie [][]int) {
	for riga := 0; riga < len(s); riga++ {
		for col := 0; col < len(s[riga]); col++ {
			if s[riga][col] == 1 {
				coppie = append(coppie, []int{riga, col})
			}
		}
	}
	return
}

func main() {
	sl := CreaSliceBidimensionale(LeggiNumero())

	InizializzaSliceBidimensionale(sl)

	coppie := Coppie(sl)

	for riga := 0; riga < len(coppie); riga++ {
		fmt.Println(coppie[riga])
	}
}
