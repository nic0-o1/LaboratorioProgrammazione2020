package main

import "fmt"

func main(){
	var tot,inserito int
	var somma,max,min,magg,mino,null int

	fmt.Print("Inserirsci un numero ")
	fmt.Scan(&tot)

	fmt.Print("Inserici ",tot," numeri ")
	for i:=0;i<tot;i++{

		fmt.Scan(&inserito)

		somma += inserito

		if(inserito>max){
			max = inserito
		}
		if(inserito<min){
			min=inserito
		}
		if(inserito>0){
			magg++
		} else if (inserito <0){
			mino++
		}else{
			null++
		}
		
	}
	fmt.Println("Somma = ",somma)
	fmt.Println("Valore minimo = ",min)
	fmt.Println("Valore massimo = ",max)
	fmt.Println("Interi >0 = ",magg)
	fmt.Println("Interi <0 = ",mino)
	fmt.Println("Interi =0 = ",null)
}