package main

import "fmt"

// func StringheAlternate(s1, s2 string) (risultato string) {

// 	lunghezzaMassima := len(s1)
// 	if lunghezzaMassima < len(s2) {
// 		lunghezzaMassima = len(s2)
// 	}

// 	s1 += strings.Repeat("-", lunghezzaMassima-len(s1))
// 	s2 += strings.Repeat("-", lunghezzaMassima-len(s2))

// 	for i := range s1 {
// 		risultato += string(s1[i]) + string(s2[i])
// 	}

// 	return
// }

func stringheAlternate(s1 string, s2 string) (ris string) {

	if len(s1) < len(s2) {
		diff := len(s2) - len(s1)

		for i := 0; i < diff; i++ {
			s1 += "-"
		}
	} else {
		diff := len(s1) - len(s2)

		for i := 0; i < diff; i++ {
			s2 += "-"
		}
	}
	for i := range s1 {
		ris += string(s1[i]) + string(s2[i])
	}
	return
}

func main() {
	var s1, s2 string

	fmt.Scan(&s1)
	fmt.Scan(&s2)

	fmt.Println(stringheAlternate(s1, s2))
}
