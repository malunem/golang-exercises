//Scrivere un programma che legga da **standard input** le età di due persone (espresse in anni) e calcoli:

// la somma delle età inserite;
// la media delle età inserite;
// la media delle età inserite arrotondata per difetto all'intero inferiore;
// la media delle età inserite arrotondata per eccesso all'intero superiore;
// la somma e la media delle età che le due persone avranno fra 10 anni.

package main

import (
	"fmt"
	"math"
)

func main() {
	var age1, age2 int

	fmt.Print("Inserisci l'età della prima persona: ")
	fmt.Scan(&age1)

	fmt.Print("Inserisci l'età della seconda persona: ")
	fmt.Scan(&age2)

	var sum_now, avg_floor, avg_ceil, sum_10years int
	var avg_now, avg_10years float64

	sum_now = age1 + age2
	avg_now = float64(sum_now) / 2
	avg_floor = int(math.Floor(avg_now))
	avg_ceil = int(math.Ceil(avg_now))
	sum_10years = sum_now + 20
	avg_10years = float64(sum_10years) / 2

	fmt.Println("La somma delle età è:", sum_now)
	fmt.Println("La media delle età è:", avg_now)
	fmt.Println("La media delle età arrotondata per difetto è:", avg_floor)
	fmt.Println("La media delle età arrotondata per eccesso è:", avg_ceil)
	fmt.Println("La somma delle età tra 10 anni sarà:", sum_10years)
	fmt.Println("La media delle età tra 10 anni sarà:", avg_10years)

}
