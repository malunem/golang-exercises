package main

import "fmt"

func main() {
    
    var ultimo string
    
    for {
        var cognome string 
        fmt.Println("Inserisci un cognome:")
        fmt.Scanln(&cognome)
        
        if cognome == "" {
            break
        }
        
        if cognome > ultimo {
            ultimo = cognome
        }
    }
    
    fmt.Println("ultimo cognome:", ultimo)
}
