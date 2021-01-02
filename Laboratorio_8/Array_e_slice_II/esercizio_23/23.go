package main

import(
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)

func genera(soglia int) (ris []int){
	rand.Seed(int64(time.Now().Nanosecond()))

	for val := int(soglia) + 1; val > int(soglia); {
		val = rand.Intn(100)
		ris = append(ris, val)
	}
	return
}

func main(){
	soglia,err:= strconv.Atoi(os.Args[1])

	if err == nil{
		val := genera(soglia)
		fmt.Println("Valori generati:",val)

		fmt.Println("Valori sopra soglia:",val[:len(val)-1])// so che l'ultimo è quello che ha interrotto il ciclo
	}
}