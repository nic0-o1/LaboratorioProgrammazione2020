package main

import "fmt"

func main(){
	var text string

	fmt.Print("Inserisci una stringa di testo: ")

	fmt.Scan(&text)

	for _, c := range text {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
}