package main

import "fmt"

func Tabellina(numero int){
	
	if(numero>=1 && numero<=9){
		fmt.Printf("Tabellina del %d :",numero)
		for i := 0; i <= 10; i++{
			fmt.Print(" ",i*numero)
	}
		
	}
}

func main(){
	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	Tabellina(numero)
}