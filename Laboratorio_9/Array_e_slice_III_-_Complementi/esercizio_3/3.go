package main

import (
	"fmt"
	"os"
	"strconv"
)

func CreaTavolaPitagorica(n int) (tavola [][]int) {
	tavola = make([][]int, n)

	for riga := 0; riga < n; riga++ {
		tavola[riga] = make([]int, n)
		for col := 0; col < len(tavola[riga]); col++ {
			tavola[riga][col] = (riga + 1) * (col + 1)
		}
	}
	return
}

func StampaTavolaPitagorica(sl [][]int){
	for riga := 0; riga <len(sl);riga++{
		for col := 0; col <len(sl[riga]);col++{
			fmt.Printf("%4d",sl[riga][col])
		}
		fmt.Println()
	}
}

func main() {

	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		StampaTavolaPitagorica(CreaTavolaPitagorica(n))
	}
}
