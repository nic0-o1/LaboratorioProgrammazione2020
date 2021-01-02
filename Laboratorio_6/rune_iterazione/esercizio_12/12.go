package main

import "fmt"

func main(){
	var n int

	fmt.Scan(&n)

	if n<=0{
		fmt.Println("Dimensione non sufficiente")
	}else{
		var stringaSpazi, stringaAsterischi string
		
		for i:=0; i< 2*n+1; i++{
			stringaAsterischi+="*"
		}
		for i:=0; i<3*n;i++{
			stringaSpazi+=" "
		}

		for i:=0; i< 2*n+1; i++{
			fmt.Print(stringaSpazi[0:i*n])
			fmt.Print(stringaAsterischi[0:i+1])
			fmt.Print(stringaSpazi[0:i*n])
		}
	}
}