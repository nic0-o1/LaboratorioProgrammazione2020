package main

import(
	"fmt"
	"os"
	"bufio"
	"unicode"
)

func LeggiTesto() string{
	scanner := bufio.NewScanner(os.Stdin)
	var testo string
		for scanner.Scan() {
			testo += scanner.Text() + "\n"
		}

	
	return testo
}

func StatisticheParole(s string) (parole int,lun int){
	var ultimoCarattereLettera bool
	for _,v := range s{
		if unicode.IsLetter(v){
			lun++ 
			ultimoCarattereLettera = true
		}else{
			if(ultimoCarattereLettera){
				parole++
				ultimoCarattereLettera = false
			}
		}
	}
	return parole,lun
}

func main(){
	var testo string

	fmt.Println("Inserisci un testo su più righe (termina con Ctrl+D):")
	
	testo=LeggiTesto()

	par,lun := StatisticheParole(testo)

	fmt.Println("Statistiche:")
	fmt.Println("Numero di parole: ",par)
	fmt.Println("Lunghezza media: ",float64(lun)/float64(par))

}