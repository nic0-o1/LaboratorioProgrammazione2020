package main

import "fmt"

func inMaiuscolo(carattere rune) rune{
	if(carattere >='a' && carattere <='z'){
		return 'A'+(carattere-'a')
	}
	return carattere
}

func inMinuscolo(carattere rune) rune{
	if(carattere >='A' && carattere <='Z'){
		return 'a'+(carattere-'A')
	}
	return carattere
}

func main(){
	var s string

	fmt.Scan(&s)
	fmt.Print("Testo in maiuscolo: ")
	for _,v:=range(s){
		fmt.Printf("%c",inMaiuscolo(v))
	}
	fmt.Println()
	
	fmt.Print("Testo in minuscolo: ")
	for _,v:=range(s){
		fmt.Printf("%c",inMinuscolo(v))
	}
	fmt.Println()
}