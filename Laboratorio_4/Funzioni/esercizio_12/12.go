package main

import "fmt"

func StampaRigaInizioAsterisco(lunghezza int){
	for i:=0;i<lunghezza;i++{
		if(i%2 == 0 ){
			fmt.Print("*")
		}else{
			fmt.Print("+")
		}
	}
}
func StampaRigaInizioPiù(lunghezza int){
	for i:=0;i<lunghezza;i++{
		if(i%2 == 0 ){
			fmt.Print("+")
		}else{
			fmt.Print("*")
		}
	}
}
func StampaScacchiera(dimensione int){
	for i:=0;i<dimensione;i++{
		if(i%2 == 0 ){
			StampaRigaInizioAsterisco(dimensione)
		}else{
			StampaRigaInizioPiù(dimensione)
		}
		fmt.Println()
	}
}

func main(){
	var dimensione int

	fmt.Println("Inserisci la dimensione: ")
	fmt.Scan(&dimensione)

	StampaScacchiera(dimensione)
}