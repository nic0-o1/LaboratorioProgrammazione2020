package main

import "fmt"

func main() {
	//la lunghezza dell'array si basa sull'indice massimo specificato in inizializzazione
	// a => ["","","C","","E"] => 5 indice con valore massimo specificato --> lunghezza 6 
	a := [...]string{5: "E", 3: "C"}

	fmt.Printf("Tipo: %T  - Valore: %v\n", a, a)

	for i := range a {
		fmt.Println("Indice", len(a)-1-i, " - Valore:", a[len(a)-1-i])
	}
	fmt.Println()

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()

	for i := 0; i < len(a); i++ {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()

}
