package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return 
}

func TraduciCarattereInFarfallino(carattere rune) (carattereTrasformato string) {
	carattereTrasformato = string(carattere)
	if strings.ContainsRune("aeiouèéòóùúàáíì", carattere) {
		
		carattereTrasformato += "f" + carattereTrasformato
	}
	if strings.ContainsRune("AEIOUÀÁÈÉÌÍÓÒÚÙ", carattere) {
		carattereTrasformato += "F" + carattereTrasformato
	}
	return
}

func TraduciTestoInFarfallino(s string) (converted string) {
	for _, c := range s {
		converted += TraduciCarattereInFarfallino(c)
	}
	return 
}

func main(){
	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D):")
	
	testo:=LeggiTesto();

	fmt.Println("Risultato:")

	fmt.Println(TraduciTestoInFarfallino(testo))
}
