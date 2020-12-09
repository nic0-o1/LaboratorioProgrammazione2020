package main

import "fmt"

func main(){

	var primo,ultimo int

	fmt.Print("Primo numero: ")
	fmt.Scan(&primo)

	fmt.Print("Secondo numero: ")
	fmt.Scan(&ultimo)

	var somma int
	for i:= primo;i<ultimo;i++{
		if i%2==0{
			somma += i 
		}
	}

	fmt.Print(somma)
}