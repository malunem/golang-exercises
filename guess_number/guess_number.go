// Scrivere un programma che:

// 1) Genera in modo casuale un numero intero compreso nell'intervallo che va da 1 a 100, estremi inclusi (sia `numeroGenerato` la variabile intera in cui viene memorizzato il numero generato, come indicato nella consegna deve valere 1<= `numeroGenerato` <= 100).

// 2) Chiede iterativamente all'utente di inserire da **standard input** un numero intero; ad ogni iterazione il programma controlla se il numero inserito è uguale al numero memorizzato in `numeroGenerato`:

// + se sono uguali, il programma termina;

// + se sono diversi, il programma fornisce un suggerimento specificando se il numero inserito è più alto o più basso di quello memorizzato in `numeroGenerato`.

// L'output stampato a video dal programma deve essere quello riportato nell'**Esempio d'esecuzione** (eccezion fatta per i numeri inseriti dall'utente).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(100)

	counter := 1

	for {
		var attempt int
		fmt.Print("Tentativo n°", counter, ": ")
		fmt.Scan(&attempt)

		if attempt == randomNumber {
			fmt.Println("Hai indovinato! N° di tentativi:", counter)
			break
		} else if attempt < randomNumber {
			fmt.Println("Numero troppo basso! Riprova")
		} else {
			fmt.Println("Numero troppo alto! Riprova")
		}

		counter++
	}
}
