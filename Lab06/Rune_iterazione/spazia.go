package main

import "fmt"

func main() {

    var testo string
    fmt.Print("Inserisci una stringa di testo: ")
    fmt.Scan(&testo)
    
    for _, v := range testo {
        
        fmt.Printf("%c ", v)
    }
    fmt.Println()
    
}
