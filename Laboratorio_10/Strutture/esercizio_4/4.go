package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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
	Contatti []Contatto
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
		n,_ := strconv.Atoi(istr[4])
		r = InserisciContatto(r, istr[1], istr[2], istr[3], uint(n), istr[4], istr[5], istr[6])
	case "E":
		r = EliminaContatto(r, istr[1], istr[2])
	case "S":
		StampaRubrica(r)
	case "A":
		n,_:= strconv.Atoi(istr[4])
		r = AggiornaContatto(r, istr[1], istr[2], istr[3], uint(n), istr[4], istr[5],istr[6])
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
	return Rubrica{}
}

func InserisciContatto(r Rubrica, nome, cognome string, via string, numero uint, cap, citta, telefono string) Rubrica {
	for _, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			return r
		}
	}
	r.Contatti = append(r.Contatti, Contatto{Cognome: cognome, Nome: nome, Telefono: telefono, Ind: Indirizzo{via, cap, citta, numero}})
	return r
}

func EliminaContatto(r Rubrica, nome, cognome string) Rubrica {
	for i, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			r.Contatti = append(r.Contatti[:i], r.Contatti[i+1:]...)
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
	for i, v := range r.Contatti {
		fmt.Printf("[%v] - ", i+1)
		StampaContatto(v)
	}
}

func AggiornaContatto(r Rubrica, nome, cognome string, via string, numero uint, cap, citta string, telefono string) Rubrica {
	for i, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			r.Contatti[i].Telefono = telefono
			r.Contatti[i].Ind = Indirizzo{Cap: cap, Citta: citta, NumeroCivico: numero, Via: via}
		}
	}
	return r
}
