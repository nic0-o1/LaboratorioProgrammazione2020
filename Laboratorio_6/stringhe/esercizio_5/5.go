package main

import (
	"fmt"
	"strconv"
)

func main() {
	var out int
	for {
		var inpunt string

		fmt.Println("Inserisci un numero: ")
		fmt.Scan(&inpunt)

		result, err := strconv.Atoi(inpunt)

		if err == nil {
			out += result
		} else {
			break
		}
		
	}
	fmt.Println("Somma: ", out)
}
