package main

import "fmt"

func ÈPrimo(n int) bool{
	for i:=2;i<n;i++{
		if(n%i==0){
			return false
		}
	}
	return true
}

func NumeriPrimi(n int) {
	fmt.Println("Numeri primi inferiori a", n)
	var numero int
	for numero = 2; numero < n; numero++ {
		if ÈPrimo(numero) {
			fmt.Print(numero, " ")
		}
	}
	fmt.Println()
}

func main() {
	var soglia int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia > 0 {
		NumeriPrimi(soglia)
	} else {
		fmt.Println("La soglia inserita non è positiva")
	}

}