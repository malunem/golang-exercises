package main

import "fmt"

func main() {

    var soglia int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&soglia)
    
    if soglia <= 0 {
        fmt.Println("La soglia inserita non è positiva")
    } else {
        fmt.Println("Numeri primi gemelli inferiori a", soglia)
        NumeriPrimiGemelli(soglia)
    }
}

func NumeriPrimiGemelli(limite int) {

    for i := 2; i < limite-2; i++ {
        
        if èPrimo(i) && èPrimo(i+2) {
        
            fmt.Print("(", i, ",", i+2, ") ")
        }
    }
    fmt.Println()
}

func èPrimo(n int) bool {

    for i := 2; i < n; i++ {
    
        if n%i == 0 {
            return false
        }
    }
    
    return true
}
