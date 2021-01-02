package main

import (
	"fmt"
	"unicode"
)

func èVocale(l rune) bool{
	if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u'{
		return true
	}
	return false
}

func main(){
	var testo string
	var vocMai,vocMin,consMai,consMin int

	fmt.Println("Inserisci testo")

	fmt.Scan(&testo)

	for _,v := range testo {
		if unicode.IsLetter(v){
			if unicode.IsUpper(v){
				if(èVocale(unicode.ToLower(v))){
					vocMai++
				}else{
					consMai++
				}
			}else{
				if(èVocale(v)){
					vocMin++
				}else{
					consMin++
				}
			}
		}
	}
	fmt.Println("Vocali maiuscole: ",vocMai)
	fmt.Println("Consonanti maiuscole: ",consMai)
	fmt.Println("Vocali minuscole: ",vocMin)
	fmt.Println("Consonanti minuscole: ",consMin)
}