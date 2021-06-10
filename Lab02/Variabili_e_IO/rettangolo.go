package main

import "fmt"

func main() {

    var base, altezza float64
    
    fmt.Print("Inserisci la base: ")
    fmt.Scan(&base)
    
    fmt.Print("Inserisci l'altezza: ")
    fmt.Scan(&altezza)
    
    fmt.Println("Permimetro =", (base+altezza)*2)
    fmt.Println("Area =", base*altezza)
}
