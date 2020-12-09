package main

import "fmt"

func Tabellina(numero int) bool{
	
	if(numero>=1 && numero<=9){
		fmt.Printf("Tabellina del %d :",numero)
		for i := 0; i <= 10; i++{
			fmt.Print(" ",i*numero)
	}
	return true
		
	}
	return false
}

func main(){
	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	if(!Tabellina(numero)){
		fmt.Println("Programma terminato")
	}
}