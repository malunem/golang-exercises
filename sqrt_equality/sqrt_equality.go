// Scrivere un programma che:  
// 1) Legga da **standard input** un numero reale.  
// 2) Calcoli la radice quadrata del numero letto (sia `x` la variabile di tipo `float64` in cui è memorizzato il valore; la radice quadrata di del valore reale può essere calcolata utilizzando la funzione `math.Sqrt` del package `math` nel seguente modo `radiceQuadrata := math.Sqrt(x)`).  
// 3) Stampi a video `x, "uguale a", radiceQuadrata, "*", radiceQuadrata, "\n"` nel caso in cui `radiceQuadrata*radiceQuadrata` sia uguale a `x` e `x, "diverso da", radiceQuadrata, "*", radiceQuadrata, "\n"` altrimenti.  

package main 

import (
	"fmt"
	"math"
)

func main() {
	num := askRealNumber()

	square_root := math.Sqrt(num)

	if sqrtEqualsNum(square_root, num) {
		fmt.Printf("%f uguale a %f * %f\n", num, square_root, square_root)
	} else {
		fmt.Printf("%f diverso da %f * %f\n", num, square_root, square_root)
	}
}

func askRealNumber() (num float64) {
	fmt.Print("Inserisci un numero reale: ")
	fmt.Scan(&num)

	return num
}

func sqrtEqualsNum(square_root float64, num float64) (bool) {
	return square_root*square_root == num 
}