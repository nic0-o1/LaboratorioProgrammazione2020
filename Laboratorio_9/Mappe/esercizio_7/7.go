package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	parole := Conta(SeparaParole(LeggiTesto()))
	fmt.Println(parole)
	fmt.Println("Parole distinte: ", len(parole))
	for k, v := range parole {
		fmt.Printf("%s: %d\n", k,v)
	}

	// fmt.Println(Conta(SeparaParole(LeggiTesto())))

}
func Conta(sl []string) (ripetizioni map[string]int) {
	ripetizioni = make(map[string]int)
	for _, v := range sl {
		ripetizioni[v]++
	}
	return
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func SeparaParole(s string) (parole []string) {
	testo := ""
	for _, v := range s {
		if unicode.IsLetter(v) {
			testo += string(v)
		} else if testo != "" {
			parole = append(parole, testo)
			testo = ""
		}
	}
	return
}
