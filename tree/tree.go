// # Albero

// Scrivere un programma che legga da **standard input** un intero `n` e, come mostrato nell'**Esempio d'esecuzione**, stampi a video un albero utilizzando il carattere `*` (asterisco) per rappresentarne la chioma ed il carattere `+` (più) per rappresentarne il tronco:
// * La chioma dell'albero deve essere alta `n` righe e, nel punto di larghezza massima, larga `2*n-1` colonne.
// * Il tronco dell'albero deve essere rappresentato con un quadrato di altezza e larghezza pari a `3` caratteri.
// * Il tronco dell'albero deve essere centrato rispetto alla chioma dell'albero.

// Se `n<=0`, il programma stampa solo il tronco dell'albero.

// ##### Esempio d'esecuzione:

// ```text
// $ go run albero.go
// 5
//       *
//      ***
//     *****
//    *******
//   *********
//      +++
//      +++
//      +++

package main

import "fmt"

func main() {
	treeHeight := askSize()
	treeLength := (2 * treeHeight) - 1
	treeCenter := treeHeight - 1
	printCanopy(treeHeight, treeLength, treeCenter)
	printTrunk(treeHeight, treeLength, treeCenter)
}

func askSize() (height int) {
	fmt.Print("inserisci un numero intero: ")
	fmt.Scan(&height)

	if height < 0 {
		height = 0
	}

	return height
}

func printCanopy(height, length, center int) {

	//rows
	for i := 0; i < height; i++ {

		//cols
		for j := 0; j < length; j++ {

			isCentered := (j >= center-i && j <= center+i) || center <= 0

			if isCentered {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func printTrunk(treeSize, treeLength, treeCenter int) {

	if treeLength <= 0 {
		treeLength = 3
	}

	//rows
	for i := 0; i < 3; i++ {

		//cols
		for j := 0; j < treeLength; j++ {

			isCentered := ((j >= treeCenter-1) && (j <= treeCenter+1) || treeCenter <= 0)

			if isCentered {
				fmt.Print("+")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 5
//       *
//      ***
//     *****
//    *******
//   *********
//      +++
//      +++
//      +++
