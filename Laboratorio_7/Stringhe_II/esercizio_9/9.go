package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

func LeggiTesto() (testo string ){
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		testo+=scanner.Text()+"\n"
	}
	return testo[:len(testo)]
}

func LeggiNumero() (num int){
	fmt.Scan(&num)
	return
}

func CifraCarattere(c rune, chiave int) rune{
	if(c != '\n' ||  !unicode.IsSpace(c)){
		fmt.Println("Ricevuto: ",c)
		return c+rune(chiave)
	}else{
		return c
	}
}


func CifraTesto(s string, chiave int) (testo string){
	for _,c := range(s){
		testo+=string(CifraCarattere(c,chiave))
	}
	return
}

func main(){
	fmt.Print("Inserisci un numero: ")
	chiave := LeggiNumero()

	fmt.Println("Inserisci un testo su più righe (termina con CTRL D):")

	daCifrare:= LeggiTesto()

	fmt.Println("Testo cifrato:")

	fmt.Println(CifraTesto(daCifrare,chiave))

}
