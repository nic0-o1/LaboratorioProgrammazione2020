package main

import (
	"fmt"
	"os"
	"strconv"
)

func LeggiNumeri() (numeri []int) {
	for _, v := range os.Args[1:] {
		if num, err := strconv.Atoi(v); err == nil {
			numeri = append(numeri, num)
		}
	}
	return
}

func ControllaSequenza(numeri []int) (stato bool, posozione int) {
	if len(numeri) > 0 {
		// for i := 1; i < len(numeri); i += 2 {
		// 	if numeri[i] >= numeri[i-1] {
		// 		return false, i
		// 	}
		// }
		// for i := 2; i < len(numeri); i += 2 {
		// 	if numeri[i] < numeri[i-1] {
		// 		return false, i
		// 	}
		// }

		for i := 1; i < len(numeri); i++{
			if i % 2 == 0{
				if numeri[i] < numeri[i-1] {
					return false, i
				}	
			}else{
				if numeri[i] >= numeri[i-1] {
					return false, i
				}
			}
		}

	} else {
		return false, 0
	}
	return true, -1
}

func main() {
	numeri := LeggiNumeri()

	if stato, pos := ControllaSequenza(numeri); stato {
		fmt.Println("Sequenza valida")
	} else {
		fmt.Printf("Valore in posizione %v non valido.\n", pos)
	}
}
