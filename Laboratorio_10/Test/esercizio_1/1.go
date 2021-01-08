package main

import(
	"fmt"
)
func Pari(n int) bool{
	return n % 2 == 0
}

func Dispari(n int) bool {
	return n % 2 != 0
}

func main(){
	fmt.Print("Inserisci un numero: ")
	var n int
	fmt.Scan(&n)

		if Pari(n){
			fmt.Println("In numero",n,"è pari")
		} else if Dispari(n){
			fmt.Println("In numero",n,"è dispari")
		}
	
}