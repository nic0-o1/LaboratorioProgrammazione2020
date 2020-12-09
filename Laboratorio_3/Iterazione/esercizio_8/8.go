package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	/* inizializzazione del generatore di numeri casuali */
	rand.Seed(int64(time.Now().Nanosecond())) // se non lo metto genera sempre lo stesso numero
	/* generazione di un numero casuale compreso nell'intervallo 
	che va da 0 a 99 (estremi inclusi) */
	var numeroGenerato int = rand.Intn(100)
	var input int
	var tentativi int =1 

	for{
		fmt.Print("Tentativo n° ",tentativi,": ")
		fmt.Scan(&input)

		if(input != numeroGenerato){
			if(input > numeroGenerato){
				fmt.Println("Hai inserito un numero troppo alto. Riprova")
			}else{
				fmt.Println("Hai inserito un numero troppo basso. Riprova")
			}
		}else{
			break
		}
		tentativi++
	}
	fmt.Println("Hai indovinato in ",tentativi, "tentativi")
}
