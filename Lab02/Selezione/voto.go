package main

import "fmt"

func main() {

    var v int
    fmt.Print("Inserisci il voto: ")
    fmt.Scan(&v)
    
    if v < 0 || v > 100 { 
        fmt.Println("Errore")
    } else if v < 60 {
        fmt.Println("Insufficiente")
    } else if v < 70 {
        fmt.Println("Sufficiente")
    } else if v < 80 {
        fmt.Println("Buono")
    } else if v < 90 {
        fmt.Println("Distinto")
    } else {
        fmt.Println("Ottimo")
    }
}
