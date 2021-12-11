//Scrivere un programma che legga da **standard input** una distanza in Km ed effettui la conversione di tale distanza in miglia (1 Km = 0.62137 mi).

package main

import "fmt"

func main() {

	var distance_km float64

	fmt.Print("Inserisci distanza in km: ")
	fmt.Scan(&distance_km)

	var distance_miles float64 = distance_km * 0.62137

	fmt.Println("La distanza in miglia è:", distance_miles)
}
