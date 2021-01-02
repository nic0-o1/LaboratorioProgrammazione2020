package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	StampaIstogramma(Occorrenze(LeggiTesto()))
}

func LeggiTesto() (testo string) {
	scanenr := bufio.NewScanner(os.Stdin)

	for scanenr.Scan() {
		testo += scanenr.Text()
	}
	return
}

func Occorrenze(s string) (totale map[rune]int) {
	totale = make(map[rune]int)
	for _, v := range s {
		if unicode.IsLetter(v) {
			totale[v]++
		}
	}
	return
}

func StampaIstogramma(occorrenze map[rune]int) {
	keys := []rune{}
	for k := range occorrenze {
		keys = append(keys, k)
	}
	
	SortRunes(keys)

	for _, v := range keys {
		fmt.Printf("%c: %v\n", v, strings.Repeat("*", occorrenze[v]))
	}
}

func SortRunes(a []rune) {
	for i := 0; i < len(a)-1; i++ {
		indiceMin := i
		for j := i + 1; j < len(a); j++ {
			if a[indiceMin] > a[j] {
				indiceMin = j
			}
		}
		a[i], a[indiceMin] = a[indiceMin], a[i]
	}
}
