package main

import "fmt"
import "os"
import "strconv"


func Fattoriale(n int) int{
	if n == 1 {
		return 1
	}else{
		return n*Fattoriale(n-1)
	}
}

func Fattoriali(n int) (f []int){
	f = make([]int,n)

	for i:=0;i<n;i++{
		f[i]= Fattoriale(i+1)
	}
	return
}

func main(){
	num,err := strconv.Atoi(os.Args[1])

	if err == nil {
		if num > 0{
			fmt.Println("Fattoriali:",Fattoriali(num))
		}
	}
}