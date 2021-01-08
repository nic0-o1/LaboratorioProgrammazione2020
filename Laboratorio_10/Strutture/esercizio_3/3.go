package main

import (
	"fmt"
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

func NuovaRubrica() Rubrica {
	return Rubrica{}
}

func InserisciContatto(r Rubrica, nome,cognome string, via string, numero uint, cap, citta, telefono string) Rubrica {
	for _, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			return r
		}
	}
	r.Contatti =append(r.Contatti, Contatto{Cognome: cognome, Nome: nome, Telefono: telefono, Ind: Indirizzo{via, cap, citta, numero}})
	return r
}

func EliminaContatto(r Rubrica, nome,cognome string) Rubrica {
	for i, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			r.Contatti = append(r.Contatti[:i],r.Contatti[i+1:]...)
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

func AggiornaContatto(r Rubrica, nome,cognome string, via string, numero uint, cap, citta string, telefono string) Rubrica {
	for i, v := range r.Contatti {
		if v.Nome == nome && v.Cognome == cognome {
			r.Contatti[i].Telefono = telefono
			r.Contatti[i].Ind = Indirizzo{Cap: cap, Citta: citta, NumeroCivico: numero, Via: via}
		}
	}
	return r
}
func main() {
	rubrica := NuovaRubrica()

	rubrica = InserisciContatto(rubrica, "Mario", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")
	rubrica = InserisciContatto(rubrica, "Mario", "Rossi",
		"Via Festa del Perdono", 11, "", "", "")
	rubrica = InserisciContatto(rubrica, "Anna", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")
	rubrica = InserisciContatto(rubrica, "Carlo", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")

	StampaRubrica(rubrica)

	rubrica = EliminaContatto(rubrica, "Mario", "Rossi")
	
	rubrica = EliminaContatto(rubrica, "Carlo", "Verdi")
	StampaRubrica(rubrica)


	rubrica = AggiornaContatto(rubrica, "Anna", "Rossi", "Via S. Sofia", 25, "20122", "Milano", "")

	StampaRubrica(rubrica)
}
