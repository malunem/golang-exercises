// Scrivere un programma che legga da **standard input** un numero intero `n` e che, come mostrato nell'**Esempio di esecuzione**, stampi a video un quadrato di `n` righe costituite ciascuna da `n`
// simboli intervallati da spazi, alternando fra loro righe costituite solo da simboli `*` (asterisco) intervallati da spazi, righe
// costituite solo da simboli `+` (più) intervallati da spazi e righe
//   costituite solo da simboli `o` (lettera o) intervallati da spazi.

package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&size)

	printSquare(size)
}

func printSquare(size int) {

	//row
	for i := 0; i < size; i++ {

		symbol := selectSymbol(i)

		//cols
		for j := 0; j < size; j++ {
			fmt.Print(symbol)
		}

		fmt.Println()
	}
}

func selectSymbol(i int) (symbol string) {
	selector := i % 3

	if selector == 0 {
		symbol = "* "
	} else if selector == 1 {
		symbol = "+ "
	} else {
		symbol = "o "
	}

	return symbol
}
