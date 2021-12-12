// Scrivere un programma che legga legga da **standard input** un numero intero `n` e che, come mostrato nell'**Esempio di esecuzione**, stampi a video un quadrato di lato `n` in cui:
// * il perimetro è rappresentato con il carattere `*` (asterisco);
// * l'area interna è rappresentata con il carattere `+` (più);
// * i caratteri di ogni riga del quadrato sono separati tra di loro da un carattere spazio.

package main

import "fmt"

func main() {

	size := askSize()
	printSquare(size)
}

//asks square size to user
func askSize() (size int) {
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&size)

	return size
}

//prints square based on input size
func printSquare(size int) {

	//rows
	for i := 0; i < size; i++ {

		//cols
		for j := 0; j < size; j++ {

			symbol := selectSymbol(i, j, size)
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}

//selects the appropriate symbol for the coordinates in relation to square size
func selectSymbol(row, col, size int) (symbol string) {

	isBorder := row == 0 || row == size-1 || col == 0 || col == size-1

	if isBorder {
		symbol = "* "
	} else {
		symbol = "+ "
	}

	return symbol
}
