package main

import (
	"fmt"
	"os"
	"strconv"
)

const SOGLIASUFF = 60

func LeggiNumeri() (numeri []int) {
	for _, arg := range os.Args[1:] {
		if str, err := strconv.Atoi(arg); err == nil {
			numeri = append(numeri, str)
		}
	}
	return
}

func FiltraVoti(voti []int) (sufficienti, insufficienti []int) {
	for _, voto := range voti {
		if voto >= SOGLIASUFF {
			sufficienti = append(sufficienti, voto)
		} else {
			insufficienti = append(insufficienti, voto)
		}
	}
	return
}

func main() {
	voti := LeggiNumeri()

	suff, insuff := FiltraVoti(voti)

	fmt.Println("Voti sufficienti:", suff)
	fmt.Println("Voti insiffucienti:", insuff)
}
