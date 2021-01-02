package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func LeggiTesto() (date []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "" {
			return
		}
		date = append(date, scanner.Text())
	}
	return
}

func DaInvertire(s string) bool {
	//g/m/aaaa, gg/m/aaaa, g/mm/aaaa, o gg/mm/aaaa
	return s[1] == '/' || s[2] == '/' //inizia con i giorni ?
}

func Inverti(data string) (dataInvertita string) {
	token := ""
	for _, v := range data {
		fmt.Printf("Ricevuto: %v. Controllo %c TOKEN: %v DATA INVERTICA: %v\n",data,v,token,dataInvertita )
		if v != '/' {
			token += string(v)
		} else {
			dataInvertita = "/" + token + dataInvertita
			token = ""
		}
	}
	dataInvertita = token + dataInvertita

	return
}

func Formatta(data string) string {
	if DaInvertire(data) {
		data = Inverti(data)
	}

	dataRune := []rune(data)

	if len(dataRune) == 8 {
		dataRune = append(dataRune[:4], '-', '0', dataRune[5], '-', '0', dataRune[7])
	} else if len(dataRune) == 9 {
		if dataRune[6] == '/' {
			dataRune = append(dataRune[:4], '-', '0', dataRune[5], '-', dataRune[7], dataRune[8])
		} else {
			dataRune = append(dataRune[:4], '-', dataRune[5], dataRune[6], '-', '0', dataRune[8])
		}
	} else {
		dataRune[4] = '-'
		dataRune[7] = '-'
	}

	return string(dataRune)
}

func main() {
	date := LeggiTesto()

	var dateFormattate []string
	for _, data := range date {
		dateFormattate = append(dateFormattate, Formatta(data))
	}

	sort.Strings(dateFormattate)
	for _, data := range dateFormattate {
		fmt.Println(data)
	}
}
