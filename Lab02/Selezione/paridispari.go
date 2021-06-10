package main

import "fmt"

func main() {

    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    fmt.Print(n, " è ")
    
    if n%2 == 0 {
        fmt.Print("pari")
    } else {
        fmt.Print("dispari")
    }
    
    fmt.Println()
}
