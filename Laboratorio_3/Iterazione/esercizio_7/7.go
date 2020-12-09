package main

import "fmt"

func main(){
	var inserito,somma,count int

	fmt.Print("Inserisci una sequenza di numeri (interrompi con numero<=0): ")

	for fmt.Scan(&inserito); inserito > 0; fmt.Scan(&inserito) {
		somma+=inserito
		count++
	}
	fmt.Println("Media aritmetica: ",float32(somma) / float32(count))
}