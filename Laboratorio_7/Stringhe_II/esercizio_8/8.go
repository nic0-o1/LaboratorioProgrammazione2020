package main

import(
	"fmt"
	"bufio"
	"os"
	"unicode"
)

func LeggiTesto() (testo string){
	scanner:= bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		testo+=scanner.Text()+"\n"
	}
	return testo
}

func Spazia(s string) (res string) {
	for _,c := range s{
		if !unicode.IsSpace(c){
			res+=string(c)+" "
		}
		if c =='\n'{
			res+="\n"
		}
	}
	return
}

func main(){
	fmt.Println("Inserisci un testo su più righe (termina con CTRL+D):")

	testo := LeggiTesto()

	fmt.Println("Risultato")

	fmt.Println(Spazia(testo))
}