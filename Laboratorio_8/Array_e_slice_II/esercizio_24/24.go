package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func LeggiTesto() (res []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return
}

func ContieneCifre(s string) bool {
	for _, v := range s {
		if unicode.IsDigit(v) {
			return true
		}
	}
	return false
}

func FiltraTesto(sl []string) (res []string) {
	for _, v := range sl {
		if ContieneCifre(v) {
			res = append(res, v)
		}
	}
	return
}

func main() {
	testo := LeggiTesto()

	fmt.Println("Testo filtrato")

	for _,v := range FiltraTesto(testo){
		fmt.Println(v)
	}
}
