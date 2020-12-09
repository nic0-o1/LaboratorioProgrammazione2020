package main

import "fmt"

func èLetteraValida(l rune) bool {
	if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
		return true
	} else {
		return false
	}
}

func èMaiuscola(l rune) bool {
	if l >= 'A' && l <= 'Z' {
		return true
	}
	return false
}

func èVocale(l rune) bool {
	switch l {
	case 'a', 'e', 'i', 'o', 'u','A','E','I','O','U':
		return true
	default:
		return false
	}
}

func main() {

	var maius, minus, cons, voc int
	var s string

	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		if èLetteraValida(rune(s[i])) {
			if èMaiuscola(rune(s[i])) {
				maius++
			} else {
				minus++
			}
			if èVocale(rune(s[i])) {
				voc++
			} else {
				cons++
			}
		}
	}

	fmt.Println("Maiuscole: ", maius)
	fmt.Println("Minuscole: ", minus)
	fmt.Println("Vocali: ", voc)
	fmt.Println("Consonanti: ", cons)
}
