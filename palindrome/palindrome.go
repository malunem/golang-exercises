// Scrivere un programma che:
// * legga da **standard input** una stringa senza spazi;
// * stampi a video il messaggio `Palindroma` nel caso in cui la stringa letta sia palindroma e `Non palindroma` altrimenti.

// Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
// * una funzione `ÈPalindroma(s string) bool` che riceve in input un valore `string` nel parametro `n` e restituisce `true` se `s` è palindroma e `false` altrimenti.

package main

import (
	"fmt"
)

func main() {
	word := askWord()
	isPalindrome := checkPalindrome(word)
	if isPalindrome {
		fmt.Println("La parola è palindroma")
	} else {
		fmt.Println("La parola non è palindroma")
	}
}

func askWord() (word string) {
	fmt.Print("Inserisci una stringa: ")
	fmt.Scan(&word)

	return word
}

func checkPalindrome(word string) (isPalindrome bool) {

	last_char := len(word) - 1
	var reversed_word string
	for i := range word {
		reversed_word += string(word[last_char-i])
	}

	return word == reversed_word
}
