package main

import "fmt"

func main(){

	var num int

	fmt.Print("Inserisci numero: ")
	fmt.Scan(&num)

	fmt.Print("Divisori di ",num,": ")

	for i:=1;i<num;i++{
		if(num %i == 0 ){
			fmt.Print(i," ")
		}
	}
}