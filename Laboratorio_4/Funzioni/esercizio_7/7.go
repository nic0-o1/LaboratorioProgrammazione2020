package main

import (
	"fmt"
	"math")

func CalcolaArea(raggio float64) float64{
	return math.Pi*(math.Pow(raggio,2))
}

func CalcolaCirconferenza(raggio float64) float64{
	return 2*math.Pi*raggio
}

func main(){
	var raggio float64

	fmt.Print("Inserisci il raggio del cerchio: ")
	fmt.Scan(&raggio)

	fmt.Println("Area del cerchio: ",CalcolaArea(raggio));
	fmt.Println("Circonferenza del cerchio: ",CalcolaCirconferenza(raggio))
}