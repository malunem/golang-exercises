package main 

import "fmt"

func main() {

    var soglia int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&soglia)
    
    if soglia <= 0 {
            fmt.Println("La soglia inserita non è positiva.")
    } else {
            fmt.Println("Numeri amichevoli inferiori a", soglia)
            NumeriAmichevoli(soglia)
    }
}

func NumeriAmichevoli(limite int) {

    for i := 1; i < limite; i++ {
    
        for j := i+1; j < limite; j++{
        
            if SonoAmichevoli(i, j) {               
                    fmt.Print("(", i, ",", j, ") ")
            }
        }
    }
    
    fmt.Println()
}

func SonoAmichevoli(n, m int) bool {

    if SommaDivisoriPropri(n) == m && 
        SommaDivisoriPropri(m) == n {
            return true
    } else {
            return false
    }
}

func SommaDivisoriPropri(n int) (somma int) {

    for i := 1; i < n; i++ {
        
        if n%i == 0 {
                somma += i
        }
    }
    
    return somma
}
