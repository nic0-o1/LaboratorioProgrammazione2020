package main

import(
	"fmt"
	"os"
	"unicode"
	"bufio"
	"strings"
)

func LeggiTesto() (testo string){
	scanenr := bufio.NewScanner(os.Stdin)

	for scanenr.Scan(){
		testo += scanenr.Text()
	}
	return 
}

func Occorrenze(s string) (totale map[rune]int){
	totale = make(map[rune]int)
	for _,v := range s{
		if unicode.IsLetter(v){
			totale[v]++
		}
	}
	return
}

func main(){
	occ := Occorrenze(LeggiTesto())

	fmt.Println("Istogramma:")
	for k,v := range occ{
		fmt.Printf("%c: ",k)
		fmt.Print(strings.Repeat("*",v))
		fmt.Println()
	}
}