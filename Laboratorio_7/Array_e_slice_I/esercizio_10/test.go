package main

import "fmt"

func main(){
	var a,b int

	fmt.Scan(&a,&b)

	fmt.Printf("A: %d B: %d \n",a,b)

	a,b = b,a

	fmt.Printf("A: %d B: %d \n",a,b)
}