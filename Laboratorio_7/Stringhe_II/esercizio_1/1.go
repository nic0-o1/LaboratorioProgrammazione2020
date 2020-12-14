package main

import (
	"fmt"
	"os"
	"bufio"
)

func LeggiTesto() string{
	scanner:= bufio.NewScanner(os.Stdin)

	var testo string
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return testo
}

func main(){
	fmt.Println("Inserisci testo (termina con CTRL+D):")

	fmt.Print("Testo letto:\n", LeggiTesto())

}