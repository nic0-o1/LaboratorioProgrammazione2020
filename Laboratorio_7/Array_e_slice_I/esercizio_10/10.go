package main

import (
	"fmt"
)



func main(){
	var n int
	var ele []int

	fmt.Print("Inserisci un numero: ")

	fmt.Scan(&n)

	ele = make([]int , n)

	fmt.Printf("Inserisci %d numeri:\n", n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&ele[i])
	}
	
	for i := range ele {
		fmt.Print(ele[len(ele)-1-i], " ")
	}
	
	
}