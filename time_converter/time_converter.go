// Scrivere un unico programma che:
// - legga da **standard input** un valore intero che specifica il tipo di conversione da effettuare:
// 1: secondi (inseriti dall’utente) in ore
// 2: secondi inseriti dall’utente in minuti
// 3: minuti inseriti dall’utente in ore
// 4: minuti inseriti dall’utente in secondi
// 5: ore inserite dall’utente in secondi
// 6: ore inserite dall’utente in minuti
// 7: minuti inseriti dall’utente in giorni e ore
// 8: minuti inseriti dall’utente in anni e giorni

//     gestendo l’insertimento di un valore di scelta non compreso tra 1 e 8;

// - legga da **standard input** un valore reale da convertire;

// - stampi a video il valore convertito.

package main

import (
	"fmt"
)

func main() {

	fmt.Print(`Scegli la conversione:
	1) secondi -> ore
	2) secondi -> minuti
	3) minuti -> ore
	4) minuti -> secondi
	5) ore -> secondi
	6) ore -> minuti
	7) minuti -> giorni e ore
	8) minuti -> anni e giorni`, "\n")

	option, value := scanOption()

	switch option {
	case 1:
		fmt.Printf("%d secondi corrispondono a %f ore\n", value, float64(value/60/60))
	case 2:
		fmt.Printf("%d secondi corrispondono a %f minuti\n", value, float64(value/60))
	case 3:
		fmt.Printf("%d minuti corrispondono a %f ore\n", value, float64(value/60))
	case 4:
		fmt.Printf("%d minuti corrispondono a %d secondi\n", value, value*60)
	case 5:
		fmt.Printf("%d ore corrispondono a %d secondi\n", value, value*60*60)
	case 6:
		fmt.Printf("%d ore corrispondono a %d minuti\n", value, value*60)
	case 7:
		days := value / 60 / 24
		hours := (value / 60) % (days * 24)
		fmt.Printf("%d minuti corrispondono a %d giorni e %d ore\n", value, days, hours)
	case 8:
		years := value / 60 / 24 / 365
		days := (value / 60 / 24) % (years * 365)
		fmt.Printf("%d minuti corrispondono a %d anni e %d giorni\n", value, years, days)
	}

}

func scanOption() (option, value int) {

	for {
		fmt.Scan(&option)

		if option >= 1 && option <= 8 {
			fmt.Print("Inserisci il valore da convertire: ")
			fmt.Scan(&value)
			break
		} else {
			fmt.Println("Inserire un numero da 1 a 8")
		}
	}

	return option, value
}
