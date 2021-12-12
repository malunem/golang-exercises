// Scrivere un programma che, dopo aver richiesto all'utente di inserire da **standard input** un numero intero `n`, stampi a video la corrispondente tabellina (moltiplicando `n` per i numeri naturali da 1 a 10)

package main

import "fmt"

func main() {
	var number int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&number)

	var i int
	for i = 1; i <= 10; i++ {
		var multiple int = number * i
		fmt.Println(number, "x", i, "=", multiple)
	}
}
