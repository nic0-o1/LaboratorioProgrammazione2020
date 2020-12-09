package main

import "fmt"

func main(){

	var num int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)

	for i:=0;i<num;i++{
		for j:=0;j<num;j++{
			//controllo righe pari o dispari
			if i%3 ==0{
				
			}
			if i%2 == 0{
				fmt.Print("* ")
			}else{
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}
}