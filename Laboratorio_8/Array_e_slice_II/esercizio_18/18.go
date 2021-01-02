package main

import (
	"bufio"
	"fmt"
	"os"
)

func LeggiTesto() []string {
	var testo []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}
		testo = append(testo, txt)
	}
	return testo
}

func Formatta(s string) string {

	date := []rune(s)

	if len(date) == 8 {
		date = append(date[:4], '-', '0', date[5], '-', '0', date[7])
	} else if len(date) == 9 {
		if date[6] == '/' {
			date = append(date[:4], '-', '0', date[5], '-', date[7], date[8])
		} else {
			date = append(date[:4], '-', date[5], date[6], '-', '0', date[8])
		}
	} else {
		date[4] = '-'
		date[7] = '-'
	}
	return string(date)
}

func main() {
	input := LeggiTesto()

	for _, v := range input {
		fmt.Println(Formatta(v))
	}
}
