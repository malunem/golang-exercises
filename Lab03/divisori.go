package main

import "fmt"

func main() {
    
    var n int
    
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&n)
    
    fmt.Print("Divisori di ", n, ": ")
    
    for i := 1; i < n-1; i++ {
    
        if n%i == 0 {
            
            fmt.Print(i, " ")
        }
    }
    fmt.Println()
}
