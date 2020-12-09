package main 

import "fmt"

func main(){
	
	var num int
	
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&num)

	for i:=0;i<num;i++{
		for j:=0;j<num;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}