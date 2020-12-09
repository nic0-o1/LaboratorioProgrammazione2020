package main

import "fmt"

func main(){
	var input int

	fmt.Scan(&input)

	for i:=1;i<=input;i++{
		if(i%3 ==0){
			fmt.Print("Fizz ")
		}else if (i%5==0){
			fmt.Print("Buzz ")
		}else if(i%3 == 0 && i%5 ==0){
			fmt.Print("Fizz buzz ")
		}else {
			fmt.Print(i," ")
		}
	}
}