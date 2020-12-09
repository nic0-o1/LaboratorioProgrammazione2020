/* si risolve il problema della scan
scan e scanln sin interrompono nel momento in cui leggono uno spazio.package Scanner

ES: 
input --> Inserisci una riga di testo: Ciao, come stai?
output --> Ciao,

è possibile ripetere la scan diverse volte all'interno di un ciclo aggiungendo manualmente uno spazio
la soluzione ha un problema --> come interrompere la lettura ?

Si utilizza uno scanner

*/

package main 

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	var testo string

	fmt.Print("Inserisci una riga di testo: ")

	scanner:= bufio.NewScanner(os.Stdin)

	if(scanner.Scan()){  // se la lettura va a buon fine ritorna true --> termina con CTRL + D
		testo = scanner.Text()
		fmt.Print(testo)
	}
}