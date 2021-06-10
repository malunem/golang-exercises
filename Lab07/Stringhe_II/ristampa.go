package main

import (
        "fmt"
        "os"
        "bufio"
)

func main() {
    
    var testo string
    fmt.Println("Inserisci testo (termina con CTRl+D):")
    
    testo = LeggiTesto()
    
    fmt.Printf("Testo letto:\n%s\n", testo)
}

func LeggiTesto() string {
    
    var testo string
    
    var riga *bufio.Scanner
    
    riga = bufio.NewScanner(os.Stdin) 
    
    for riga.Scan() {
        
        testo += riga.Text() + "\n"
    }
    
    return testo
}
