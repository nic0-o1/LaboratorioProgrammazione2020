package main

import (
	"fmt"
	"os"
	"strconv"
)

func Minimo(numeri []int) int {
	min := numeri[0]
	for _, n := range numeri {
		if min > n {
			min = n
		}
	}
	return min
}

func Massimo(sl []int) (max int){
	max = sl[0]
	for _,v:=range sl{
		if(v>max){
			max=v
		}
	}
	return
}

func Media(sl []int) (media float64){
	for _,v:=range sl{
		media+=float64(v)
	}
	return media/float64(len(sl))
}

func main(){

	val:= make([]int, len(os.Args)-1)

	for i,v := range os.Args[1:]{
		ris,err := strconv.Atoi(v)

		if err == nil{
			val[i]=ris
		}
	}

	fmt.Println("Minimo: ",Minimo(val))
	fmt.Println("Massimo: ",Massimo(val))
	fmt.Printf("Valore medio: %2.f",Media(val))
}