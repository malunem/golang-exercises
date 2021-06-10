package main

import (
        "fmt"
        "unicode"
        "os"
        "bufio"
)

func main() {

    //Legge testo su più righe anche vuote fino a EOF
    testo := LeggiTesto()
    
    testo = FiltraTesto(testo)
    //Stampa le righe in cui compaiono cifre
    fmt.Println("Testo filtrato:")
    
    for _, riga := range testo {
        fmt.Println(riga)
    }
}

func LeggiTesto() []string {
    //Una riga è un elemento della slice
    scanner := bufio.NewScanner(os.Stdin)
    
    var testo []string 
    
    for scanner.Scan() {
        
        testo = append(testo, scanner.Text())
    }
    
    return testo
}

func ContieneCifre(s string) bool {
    //Return true se la riga contiene una cifra
    //Usare unicode.IsDigit
    
    for _, w := range s {
            
        if unicode.IsDigit(w) {
             return true
        }
    }
    
    return false
}

func FiltraTesto(sl []string) []string {
    //Crea una slice con le righe che contengono cifre
    
    var righeConNumeri []string
    
    for _, v := range sl {
        
        if ContieneCifre(v) {
           righeConNumeri = append(righeConNumeri, v)
        }
    }
    
    return righeConNumeri
}
