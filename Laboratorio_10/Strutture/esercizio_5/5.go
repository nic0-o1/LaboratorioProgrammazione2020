package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Indirizzo struct {
	Via, Cap, Citta string
	NumeroCivico    uint
}

type Contatto struct {
	Cognome, Nome, Telefono string
	Ind                     Indirizzo
}

type Rubrica struct {
	Contatti map[string]Contatto
}

func main() {
	istruzioni := strings.Split(LeggiTesto(), "\n")
	rubrica := NuovaRubrica()

	for _, comando := range istruzioni {
		rubrica = InterpretaComando(comando, rubrica)
	}
}

func InterpretaComando(comando string, r Rubrica) Rubrica {
	istr := strings.Split(comando, ";")
	switch istr[0] {
	case "I":
		n, _ := strconv.Atoi(istr[4])
		r = InserisciContatto(r, istr[1], istr[2], istr[3], uint(n), istr[5], istr[6], istr[7])
	case "E":
		r = EliminaContatto(r, istr[1], istr[2])
	case "S":
		StampaRubrica(r)
	case "A":
		n, _ := strconv.Atoi(istr[4])
		r = AggiornaContatto(r, istr[1], istr[2], istr[3], uint(n), istr[5], istr[6], istr[7])
	}
	return r
}

func LeggiTesto() (testo string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func NuovaRubrica() Rubrica {
	return Rubrica{Contatti: make(map[string]Contatto)}
}

func InserisciContatto(r Rubrica, nome, cognome string, via string, numero uint, cap, citta, telefono string) Rubrica {
	for _, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			return r
		}
	}
	fmt.Println("Add: ", telefono)
	r.Contatti[nome+" "+cognome] = Contatto{Cognome: cognome, Nome: nome, Telefono: telefono, Ind: Indirizzo{via, cap, citta, numero}}
	return r
}

func EliminaContatto(r Rubrica, nome, cognome string) Rubrica {
	for _, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			delete(r.Contatti, nome+" "+cognome)
		}
	}
	return r

}

func StampaContatto(contatto Contatto) {
	fmt.Printf("%s %s: %s %v, %s, %s - Tel. %s\n", contatto.Nome, contatto.Cognome, contatto.Ind.Via,
		contatto.Ind.NumeroCivico, contatto.Ind.Citta, contatto.Ind.Cap, contatto.Telefono)
}

func StampaRubrica(r Rubrica) {
	fmt.Println("Rubrica:")
	count := 1
	for _, v := range r.Contatti {
		fmt.Printf("[%v] - ", count)
		count++
		StampaContatto(v)
	}
}

func AggiornaContatto(r Rubrica, nome, cognome string, via string, numero uint, cap, citta string, telefono string) Rubrica {
	for k, _ := range r.Contatti {
		if k == nome+" "+cognome {
			r.Contatti[nome+" "+cognome] = Contatto{Cognome: cognome, Nome: nome, Telefono: telefono, Ind: Indirizzo{via, cap, citta, numero}}
		}
	}
	return r
}
