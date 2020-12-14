package main

import (
	"fmt"
	"unicode"
	"strconv"
	"bufio"
	"os"
)

func LeggiTesto() (testo string){
	scanner := bufio.NewScanner(os.Stdin)

	if(scanner.Scan()){
		testo = scanner.Text()
		return testo
	}
	return ""
}

func NumeroNascosto(testo string) (res int ,er error){
	var tot string
	for _,c := range testo{
		if(unicode.IsDigit(c)){
			tot+=string(c)

		}
	}
	
	return strconv.Atoi(tot)
}

func main(){
	testo := LeggiTesto()

	res,err :=NumeroNascosto(testo)

	if (err == nil){
		fmt.Printf("Doppio del numero nascosto: %d (%d * 2)",res*2,res)
	}else{
		fmt.Println("nessun numero nascosto")
	}
}