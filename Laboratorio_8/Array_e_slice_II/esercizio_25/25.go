package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func LeggiNumero() string {
	return os.Args[1]
}

func GeneraMazzo() (mazzo string) {
	const asso = '\U0001F0B1'
	for i := 0; i < 10; i++ {
		mazzo += string(asso + i)
	}
	return
}

func EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) {
	mazzoRune := []rune(mazzo)
	indice := rand.Intn(len(mazzoRune))
	cartaEstratta = mazzoRune[indice]
	mazzoRune = append(mazzoRune[:indice], mazzoRune[indice+1:]...)
	mazzoResiduo = string(mazzoRune)

	return
}

func EstraiCarte(mazzo string, n int) {
	for i := 0; i < n; i++ {
		var cartaEstratta rune
		cartaEstratta, mazzo = EstraiCarta(mazzo)

		fmt.Printf("Estratta carta %c", cartaEstratta)
		fmt.Println("- Carte rimaste nel mazzo: ", mazzo)

	}
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	if n, err := strconv.Atoi(LeggiNumero()); err == nil {
		EstraiCarte(GeneraMazzo(), n)
	}
}
