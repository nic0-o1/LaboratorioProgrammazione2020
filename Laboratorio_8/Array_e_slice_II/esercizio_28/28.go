package main

import (
	"fmt"
	"unicode"
)

func Ordinata(sequenza string) bool {
	seq := []rune(sequenza)

	for i := 0; i < len(seq)-1; i++ {
		if seq[i] < seq[i+1] {
			return false
		}
	}
	return true
}

func EstraiNumeri(sequenza string) (sequnzaNumerata string) {
	for _, v := range sequenza {
		if unicode.IsDigit(v) {
			sequnzaNumerata += string(v)
		}
	}
	return
}

func main() {
	var sequenza string

	fmt.Scan(&sequenza)

	if Ordinata(EstraiNumeri(sequenza)) {
		fmt.Println("Sequenza nascosta ordinata")
	} else {
		fmt.Println("Sequenza nascosta non ordinata")
	}
}
