package main

import(
	"fmt"
	"bufio"
	"os"
)

func LeggiTesto() (testo string){
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		if(scanner.Text()==""){
			break
		}
		testo+=scanner.Text()+"\n"
	}
	return testo[:len(testo)-1]  // rimuovo /n
}

func InvertiStringa(s string) (reverse string){
	for i:=0;i<len(s);i++{
		reverse+=string(s[len(s)-1-i])
	}
	return
}

func main(){
	fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")

	testo:= LeggiTesto()

	fmt.Println("Testo invertito:")
	fmt.Println(InvertiStringa(testo))


}