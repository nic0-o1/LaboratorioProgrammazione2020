package main

import "fmt"

func ÈPalindroma(s string) bool{
	for i := range s{
		if(s[i] != s[len(s)-i-1]){
			return false
		}
	}
	return true
}

func main(){
	var toCheck string

	fmt.Scan(&toCheck)

	if ÈPalindroma(toCheck){
		fmt.Println("Palindroma")
	}else {
		fmt.Println("Non palindroma")
	}
}