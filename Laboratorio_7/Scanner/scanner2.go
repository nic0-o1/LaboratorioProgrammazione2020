/*
è possibile utilizzare la scan per leggere da input su più righe --> faccio la scan più volte e termino se vuoto o CTRL + D

*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var testo string
	fmt.Println("Inserisci del testo su più righe " +
		"(termina con CTRL+D o riga vuota): ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo = scanner.Text()
		if testo == "" {
			break
		}
		fmt.Println(testo)
	}
}