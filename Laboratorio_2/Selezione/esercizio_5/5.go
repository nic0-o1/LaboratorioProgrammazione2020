package main

import "fmt"

func main(){
	var input int 
	
	fmt.Print("Inserisci numero ")
	fmt.Scan(&input)

	if input <0 {
		fmt.Println(input)
	}else if input >0{
		fmt.Println("+",input)
	}else{
		fmt.Println(input)
	}	
}