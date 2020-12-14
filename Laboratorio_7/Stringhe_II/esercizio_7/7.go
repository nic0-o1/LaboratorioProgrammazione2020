package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"unicode"
)

func LeggiTesto() (testo string){
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		if scanner.Text() ==""{
			break
		}
		testo+=scanner.Text() +"\n"
	}
	return testo[:len(testo)-1] //rimosso ultimo /n
}


func TrasformaCarattere(c rune) rune {
	if (strings.ContainsRune("aeiouèéòóùúàáíì", unicode.ToLower(c))){
		if unicode.IsLower(c){
			return 'u'
		}else{
			return 'U'
		}
	}else{
		return c
	}

}

func Garibaldi(s string) (testo string){
	for _,c := range s{
		testo += string(TrasformaCarattere(c))
	}
	return 
}

func main(){
	fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")

	testo := LeggiTesto()

	fmt.Println("Risultato operazione:")

	fmt.Println(Garibaldi(testo))
}
