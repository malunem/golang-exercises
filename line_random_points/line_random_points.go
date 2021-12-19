// Scrivere un programma che legga da **standard input** un valore intero e due valori reali:
// - il primo valore è il seme (seed) `s` da utilizzare per inizializzare il generatore di numeri casuali;
// - il secondo ed il terzo valore sono il coefficiente angolare `m` e il termine noto `q` di una retta `r: y = m*x + q` sul piano cartesiano.

// Una volta terminata la fase di lettura, il programma deve generare per 10 volte una coppia di valori reali `px` e `py` che rappresentano rispettivamente l'ascissa e l'ordinata di un punto sul piano cartesiano e, per ciascun punto:
//  1. determinare se, a meno di una costante `EPSILON = 1.0e-9`, il punto sta sopra, sotto, o appartiene alla retta `r`;
//  2. stampare a video il relativo messaggio (come mostrato nell'**Esempio di esecuzione**).

// I valori `px` e `py` devono essere compresi nell'intervallo [0, 10.0).

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const EPSILON float64 = 1.0e-9

	seed := askSeed()
	m, q := askMQ()

	rand.Seed(seed)

	for i := 0; i < 10; i++ {
		px, py := getRandomPoints()

		if pointIsUnderLine(px, py, m, q, EPSILON) {
			fmt.Printf("Punto (%f, %f) - Il punto sta sotto la retta\n", px, py)
		} else if pointIsAboveLine(px, py, m, q, EPSILON) {
			fmt.Printf("Punto (%f, %f) - Il punto sta sopra la retta\n", px, py)
		} else {
			fmt.Printf("Punto (%f, %f) - Il punto appartiene alla retta\n", px, py)
		}
	}
}

func askSeed() (num int64) {
	fmt.Print("Inserisci un numero intero: ")
	fmt.Scan(&num)

	return num
}

func askMQ() (m, q float64) {
	fmt.Print("Inserisci due numeri reali m e q: ")
	fmt.Scan(&m, &q)

	return m, q
}

func getRandomPoints() (px, py float64) {
	px = rand.Float64() * 10
	py = rand.Float64() * 10

	return px, py
}

func pointIsUnderLine(px, py, m, q, EPSILON float64) bool {
	line_y := (m * px) + q

	return (line_y - py) > EPSILON
}

func pointIsAboveLine(px, py, m, q, EPSILON float64) bool {
	line_y := (m * px) + q

	return (py - line_y) > EPSILON
}
