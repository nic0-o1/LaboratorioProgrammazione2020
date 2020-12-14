package main

import(
	"fmt"
	"strconv"
	"os"
)

func LeggiNumeri() (numeri []int){
	for _,v := range os.Args[1:]{
		if res,err :=strconv.Atoi(v); err == nil{
			numeri= append(numeri,res)
		}
	}
	return
}
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

	val := LeggiNumeri()
	fmt.Println("Minimo: ",Minimo(val))
	fmt.Println("Massimo: ",Massimo(val))
	fmt.Printf("Valore medio: %.2f",Media(val))
	
}
