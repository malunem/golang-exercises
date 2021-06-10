package main

import "fmt"

func main() {

    var testo string
    
    for {
        
        var riga string 
        
        fmt.Println("Inserisci una stringa:")
        fmt.Scanln(&riga)
        
        testo += riga + " "
        
        if riga == "" {
            break
        }
    }
    
    fmt.Println(testo)
}
