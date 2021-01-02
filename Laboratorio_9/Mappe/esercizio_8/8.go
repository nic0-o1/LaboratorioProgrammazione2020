package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	saldoIniziale, operazioni := LeggiTesto()

	var operazioniFormattate []string

	sort.Strings(operazioni)

	for _, v := range operazioni {
		operazioniFormattate = append(operazioniFormattate, FormattaOperazione(v))
	}

	operazioniRaggruppate := make(map[string][]string)

	for _, operazione := range operazioniFormattate {
		data := operazione[:10]
		operaz := operazione[11:]

		operazioniRaggruppate[data] = append(operazioniRaggruppate[data], operaz)
	}
	StampaEstrattoConto(saldoIniziale, operazioniRaggruppate)

}

func StampaEstrattoConto(saldoIniziale float64, operazioni map[string][]string) {
	date := []string{}
	for k := range operazioni {
		date = append(date, k)
	}

	fmt.Printf("SALDO INIZIALE:\t\t\t %.2f\n\n", saldoIniziale)
	saldoFinale := saldoIniziale

	for _, v := range date {
		fmt.Println("DATA: ", v)
		saldoGiornaliero := 0.00
		for _, ope := range operazioni[v] {
			azione := strings.Split(ope, ";")
			if azione[0] == "P" {
				s, _ := strconv.ParseFloat(azione[1], 64)
				// fmt.Printf("Prelievo:\t\t\t %.2f\n",s)
				fmt.Printf("%-25s%11.2f\n", "Prelievo:", s)
				saldoGiornaliero -= s
				saldoFinale -= s
			} else {
				s, _ := strconv.ParseFloat(azione[1], 64)
				// fmt.Printf("Versamento:\t\t\t %.2f\n",s)
				fmt.Printf("%-25s%11.2f\n", "Versamento:", s)
				saldoGiornaliero += s
				saldoFinale += s
			}
		}
		// 	fmt.Printf("\t\t\t\t")
		// 	fmt.Println(strings.Repeat("_",11))
		// 	fmt.Printf("SALDO GIORNALIERO:\t\t %.2f\n",saldoGiornaliero)
		// 	fmt.Print("\t\t\t\t",strings.Repeat("=",11),"\n\n")
		fmt.Printf("%-25s%11s\n", "", strings.Repeat("_", 11))
		fmt.Printf("%-25s%11.2f\n", strings.ToUpper("Saldo giornaliero:"), saldoGiornaliero)
		fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 11))
	}

	fmt.Printf("%-25s%11.2f\n", strings.ToUpper("Saldo finale:"), saldoFinale)
	fmt.Printf("%-25s%11s\n\n", "", strings.Repeat("=", 11))
	// }
	// fmt.Printf("SALDO FINALE:\t\t\t %.2f\n",saldoFinale)
	// fmt.Print("\t\t\t\t",strings.Repeat("=",11))
}

func FormattaOperazione(operazione string) string {
	dateIndex := 0
	for i := range operazione {
		if rune(operazione[i]) == ';' {
			dateIndex = i
			break
		}
	}
	return FormattaData(operazione[:dateIndex], '/', '-') + operazione[dateIndex:]

}

func LeggiTesto() (saldoIniziale float64, operazioni []string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	saldoIniziale, _ = strconv.ParseFloat(scanner.Text(), 64)

	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			return
		}
		operazioni = append(operazioni, riga)

	}
	return
}

func DaInvertire(data string, sep rune) bool {
	return rune(data[4]) == sep && (rune(data[6]) == sep || rune(data[7]) == sep)
}
func Inverti(data string, sep string) (invertita string) {
	dates := strings.Split(data, sep)

	dates[0], dates[2] = dates[2], dates[0]

	return dates[0] + sep + dates[1] + sep + dates[2]
}

func FormattaData(data string, vecchioSep rune, nuovoSep rune) string {
	if DaInvertire(data, vecchioSep) {
		data = Inverti(data, string(vecchioSep))
	}
	dataRune := []rune(data)

	if len(dataRune) == 8 {
		dataRune = append(dataRune[:4], nuovoSep, '0', dataRune[5], '-', '0', dataRune[7])
	} else if len(dataRune) == 9 {
		if dataRune[1] == vecchioSep {
			dataRune = append([]rune{'0'}, dataRune[0], nuovoSep, '0', dataRune[3], nuovoSep, dataRune[5], dataRune[6], dataRune[7], dataRune[8])
		} else {
			//aaaa-mm-0g
			dataRune = append(dataRune[:2], nuovoSep, '0', dataRune[3], nuovoSep, dataRune[5], dataRune[6], dataRune[7], dataRune[8])
		}
	} else {
		dataRune[2] = nuovoSep
		dataRune[5] = nuovoSep
	}

	return string(dataRune)

}
