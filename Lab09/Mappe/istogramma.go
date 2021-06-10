package main

import (
        "fmt"
        "bufio"
        "os"
        "unicode"
)

func main() {
    //Legge da os.Stdin con bufio un testo su più righe anche vuote fino a EOF
    testo := LeggiTesto()
    
    //Stampa ogni lettera presente nel testo con tanti * quante le occorrenze della lettera 
    occorrenzeLettere := Occorrenze(testo)
    
    for i, v := range occorrenzeLettere {
        fmt.Printf("%c: ", i)à
        for j := 0; j < v; j++ {
            fmt.Printf("*")
        }
        fmt.Println()
    }
}

func LeggiTesto() string {
    var testo string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        testo += scanner.Text() + "\n"
    }
    return testo
}

func Occorrenze(s string) map[rune]int {
    //filtra le lettere con unicode.IsLetter (case sensitive)
    occorrenze := map[rune]int{}
    for _, v := range s {
        if unicode.IsLetter(v) {
                occorrenze[v]++
        }
    }
    //restituisce il numero di occorrenze delle lettere in s
    return occorrenze
}
