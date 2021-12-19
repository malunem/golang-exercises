// Scrivere un programma che legga da **riga di comando** una sequenza di valori intervallati da caratteri di spaziatura.

// Il primo valore che definisce la sequenza (da sinistra verso destra) è in posizione 0, il secondo in posizione 1, etc.

// La sequenza è valida se:
// 1. Tutti i valori letti rappresentano dei numeri interi.
// 2. Ciascun numero che appare in una posizione *dispari* all'interno della sequenza è *minore* del numero che lo precede.
// 3. Fatta eccezione per il numero che appare in posizione 0, ciascun numero che appare in una posizione *pari* all'interno della sequenza è *maggiore* del numero che lo precede.

// Nel caso in cui la sequenza letta sia valida, il programma deve stampare:

// `Sequenza valida.`

// In caso contrario, il programma deve stampare:

// `Valore in posizione POSIZIONE non valido.`

// dove `POSIZIONE` è la posizione del primo valore che invalida la sequenza.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers, areNumbers, notNumberIndex := readOsArgs()

	if !areNumbers {
		fmt.Printf("Valore in posizione %d non valido\n", notNumberIndex)
		return
	}

	sequenceIsValid, notValidIndex := checkSequence(numbers)

	if sequenceIsValid {
		fmt.Println("Sequenza valida\n")
	} else {
		fmt.Printf("Valore in posizione %d non valido\n", notValidIndex)
	}
}

func readOsArgs() (input []int, areNumbers bool, index int) {

	areNumbers = true
	index = -1

	for i, v := range os.Args[1:] {
		num, err := strconv.Atoi(v)
		if err != nil {
			areNumbers = false
			index = i
			break
		} else {
			input = append(input, num)
		}
	}

	return input, areNumbers, index
}

func checkSequence(numbers []int) (isValid bool, notValidIndex int) {
	isValid = true
	notValidIndex = -1

	for i, v := range numbers {

		if i == 0 {
			continue
		} else if i%2 != 0 {
			if v >= numbers[i-1] {
				isValid = false
				notValidIndex = i
				break
			}
		} else {
			if v <= numbers[i-1] {
				isValid = false
				notValidIndex = i
				break
			}
		}
	}

	return isValid, notValidIndex
}
