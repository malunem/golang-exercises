package main

import "fmt"

func main() {

    var soglia int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&soglia)
    
    if soglia <= 0 {
        fmt.Println("La soglia inserita non è positiva.")
        return
    }
    
    fmt.Println("Numeri primi inferiori a ", soglia)
    NumeriPrimi(soglia)
}

func NumeriPrimi(limite int) {

    for i := 2; i < limite; i++ {
        
        if èPrimo(i) {
        
            fmt.Print(i, " ")
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
