// Scrivere un programma che legga da **standard input** 4 valori a virgola mobile:
// - i primi due valori sono il coefficiente angolare `m` e il termine noto `q` di una retta `r: y = m*x + q`
// - il terzo e il quarto valore sono le coordinate `px` e `py` di un punto `P(px,py)`

// Il programma deve determinare se il punto `P` sta sopra o sotto la retta od appartiene ad essa, e stampare a video il relativo messaggio.

package main

import (
	"fmt"
)

func main() {
	var line_m, line_q, point_x, point_y float64
	fmt.Print("Inserisci m, q, x, y: ")
	fmt.Scan(&line_m, &line_q, &point_x, &point_y)

	line_y := (line_m * point_x) + line_q

	if point_y == line_y {
		fmt.Println("Il punto appartiene alla retta")
	} else if point_y < line_y {
		fmt.Println("Il punto sta sotto alla retta")
	} else {
		fmt.Println("Il punto sta sopra alla retta")
	}
}
