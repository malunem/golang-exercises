// Scrivere un programma che legga da **standard input** due stringhe senza spazi `s1` e `s2` e stampi a video la stringa creata alternando i caratteri delle stringhe `s1` e `s2`.

// A tal fine utilizzare una funzione 'StringheAlternate(s1, s2 string) (risultato string)'

// Si assuma che le stringhe lette siano interamente definite da caratteri considerati nello standard US-ASCII.

// Se le stringhe lette non sono definite dallo stesso numero di caratteri, si deve utilizzare il carattere '-' come carattere di riempimento:

package main

import (
	"fmt"
	"strings"
)

func main() {
	string_1, string_2 := askStrings()
	pairStringsLength(&string_1, &string_2)
	alternate_string := alternateStrings(string_1, string_2)
	fmt.Println(alternate_string)
}

func askStrings() (string_1, string_2 string) {
	fmt.Print("Inserisci la prima stringa: ")
	fmt.Scan(&string_1)
	fmt.Print("Inserisci la seconda stringa: ")
	fmt.Scan(&string_2)

	return string_1, string_2
}

func pairStringsLength(string_1 *string, string_2 *string) {

	length_1 := len(*string_1)
	length_2 := len(*string_2)

	if length_1 < length_2 {
		difference := length_2 - length_1
		*string_1 += strings.Repeat("-", difference)
	} else {
		difference := length_1 - length_2
		*string_2 += strings.Repeat("-", difference)
	}
}

func alternateStrings(string_1 string, string_2 string) (result_string string) {
	for i := range string_1 {
		result_string += string(string_1[i])
		result_string += string(string_2[i])
	}

	return result_string
}
