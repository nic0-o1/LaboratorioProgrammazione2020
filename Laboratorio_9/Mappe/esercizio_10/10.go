package main

import(
	"fmt"
	"os"
)

const SEPARATORE = ' '

func main(){
	sequenza := os.Args[1:]

	seq :=GeneraSottosequenze(sequenza)
	StampaSottosequenza(seq)

}

func StampaSottosequenza(sottosequenze map[string]int){
	lunghezzaMassima := 0
	for k := range sottosequenze {
		if len(k)> lunghezzaMassima{
			lunghezzaMassima = len(k)
		}
	}

	for i := lunghezzaMassima; i>=0;i--{
		for k := range sottosequenze{
			if len(k) == i{
				fmt.Printf("%s -> Occorrenze: %v\n",k,sottosequenze[k])
			}
		}
	}
}

func GeneraSottosequenze(sequenza []string) (sottosequenze map[string]int) {
	sottosequenze = make(map[string]int)
	for i := 0; i < len(sequenza)-1; i++ {
		for j := i + 1; j < len(sequenza); j++ {
			if sequenza[i] == sequenza[j]  && len(sequenza[i:j])+1 > 2{
				var sottosequenza string
				for k := i; k < j; k++ {
					sottosequenza += sequenza[k] + string(SEPARATORE)
				}
				sottosequenza += sequenza[j]
				sottosequenze[sottosequenza]++
			}
		}
	}
	return
}