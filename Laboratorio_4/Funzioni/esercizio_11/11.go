package main

import (
	"fmt"
	"math"
)

func RadiceQuadrata(numero float64) (float64, bool){
	if(numero >=0){
		return math.Sqrt(numero),true
	}else{
		return 0,false
	}
}

func main(){
	var numero float64

	var  stato bool
	var res float64

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	res,stato = RadiceQuadrata(numero)

	if(stato){
		fmt.Println("Radice quadrata: ",res)
	}else{
		fmt.Println("Il valore inserito non appartiene al dominio della funzione")
	}
}