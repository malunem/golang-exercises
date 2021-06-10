package main

import "fmt"

func main() {

    var soglia int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&soglia)
    
    if soglia <= 0 {
            fmt.Println("La soglia inserita non è positiva.")
    } else {
            fmt.Print("Numeri abbondanti: ")
            NumeriAbbondanti(soglia)
    }
}

func NumeriAbbondanti(limite int) {

    for i := 1; i < limite; i++ {
    
        if èAbbondante(i) {
                fmt.Print(i, " ")
        }
    }
    
    fmt.Println()
}

func èAbbondante(n int) bool {
    
    if SommaDivisoriPropri(n) > n {
            return true
    } else {
            return false
    }
}

func SommaDivisoriPropri(n int) somma int {

    for i := 1; i < n; i++ {
        
        if n%i == 0 {
                somma += i
        }
    }
    
    return somma
}
