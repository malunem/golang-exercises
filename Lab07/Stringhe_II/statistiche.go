package main

import (
        "fmt"
        "bufio"
        "os"
        "unicode"
)

func main() {

    var testo string
    
    fmt.Println("Inserisci un testo su più righe (termina con Ctrl+D):")
    
    testo = LeggiTesto() 
    
    parole, lettere := StatisticheParole(testo)
    
    fmt.Println("Statistiche:")
    fmt.Printf("Numero parole: %d\n", parole)
    fmt.Printf("Lunghezza media: %f\n", float64(lettere)/float64(parole))
}

func LeggiTesto() (testo string) {

    var riga *bufio.Scanner
    
    riga = bufio.NewScanner(os.Stdin)
    
    for riga.Scan() {
        
        testo += riga.Text() + "\n"
    }
    
    return testo
}

func StatisticheParole(s string) (int, int) {
    
    var parole, lettere int
    
    var lastChar bool = false
    
    for _, char := range s {
    
        if unicode.IsLetter(char) {
        
            lettere += 1
            
            if !lastChar {
            
                parole += 1
                lastChar = true
            }
            
        } else if !unicode.IsLetter(char) {
            
            lastChar = false
        }                                 
    }
    
    return parole, lettere
}
