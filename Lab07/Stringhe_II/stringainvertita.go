package main

import (
        "fmt"
        "bufio"
        "os"
)

func main() {
    
    fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")
    fmt.Printf("Testo invertito:\n%s", InvertiStringa(LeggiTesto()))
    
}

func LeggiTesto() string {

    var testo string
    var scanner *bufio.Scanner
    
    scanner = bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
    
        if scanner.Text() == "" {
            return testo
        }
        
        testo += "\n" + scanner.Text()
    }
    
    return testo
}

func InvertiStringa(s string) string {

    var testoInvertito string
    
    for i, _ := range s {
        
        testoInvertito += string(s[len(s)-1-i])
    }
    
    return testoInvertito
}
