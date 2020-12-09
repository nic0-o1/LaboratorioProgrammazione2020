package main 

import "fmt"

func main(){ 
	var s string

	fmt.Scan(&s)

	for i,_ :=range(s){
		fmt.Println("Sottostringa ",i+1,": ",s[i:(len(s)-i)])
	}
}