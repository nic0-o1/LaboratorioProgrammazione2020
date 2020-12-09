package main

import "fmt"

func main() {

	var toPrint string

	for {
		var temp string
		fmt.Println("inserisci una stringa")
		fmt.Scanln(&temp)

		if temp == "" {
			break
		} else {
			toPrint += temp + " "
		}
	}
	fmt.Println("Stringa completa:", toPrint)
}
