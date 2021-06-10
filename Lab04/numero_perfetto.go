package main

import "fmt"

func main() {

    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    if èPerfetto(n) {
        fmt.Println(n, "è perfetto")
    } else {
        fmt.Println(n, "non è perfetto")
    }
}

func èPerfetto(n int) bool {

    if SommaDivisoriPropri(n) == n {
        return true
    } else {
        return false
    }
}

func SommaDivisoriPropri(n int) int {

    var somma int
    for i := 1; i < n; i++ {
        
        if n%i == 0 {
            somma += i
        }
    }
    
    return somma
}
