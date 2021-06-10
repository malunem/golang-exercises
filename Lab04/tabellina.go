package main

import "fmt"

func main() {

	var n int

	for {
		fmt.Print("Inserisci un numero: ")

		fmt.Scan(&n)

		if !Tabellina(n) {

			fmt.Println("Programma terminato.")
			break
		}

	}
}

func Tabellina(n int) bool {

	if n < 1 || n > 9 {

		return false
	}

	fmt.Println("Tabellina del", n, ": ")

	for i := 0; i <= 10; i++ {

		fmt.Print(i*n, " ")
	}

	fmt.Println()
	return true
}
