package main

import (
	"fmt"
	"math/rand"
	"time"
)

const FACCIEDADO = 6

func main(){
	var ngioc,turni int
	var gioc,vincitori []string

	rand.Seed(int64(time.Now().Nanosecond()))

	fmt.Print("Inserisci il numero di giocatori: ")
	fmt.Scan(&ngioc)

	gioc = make([]string, ngioc)
	fmt.Printf("\nInserisci i nomi dei %d giocatori:\n",ngioc)

	for i:=0;i<ngioc;i++{
		fmt.Scan(&gioc[i])
	}
	fmt.Printf("\nInserisci il numero di turni: ")
	fmt.Scan(&turni)

	vincitori = make([]string, turni)

	for i:=0;i<turni;i++{
		fmt.Printf("\n=== Turno %d ===\n",i+1)

		var ris []int

		ris = make([]int, ngioc)
		
		for i,v := range gioc{
			ris1:=rand.Intn(FACCIEDADO)+1
			ris2:=rand.Intn(FACCIEDADO)+1

			ris[i]=ris1+ris2
			fmt.Printf("\tGiocatore: %s, lanci di valore: %d e %d\n",v,ris1,ris2)
		}
		var max,ind int
		for i:=0;i<len(ris);i++{ 
			if(ris[i]>max){
				max=ris[i]
				ind = i
			}
		}
		vincitori[i]=gioc[ind]
		fmt.Printf("\tTurno %d, vincitore: %s\n",i+1,gioc[ind])

	}
	fmt.Println("Vittorie:")

	for i,v := range vincitori{
		fmt.Printf("Vincitore tuno %d: %s\n",i+1,v)
	}
}