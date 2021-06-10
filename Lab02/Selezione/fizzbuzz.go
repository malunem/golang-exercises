package main

import "fmt"

func main() {

    var n int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    
    if n%3 == 0 {
        fmt.Print("Fizz ")
    }
    
    if n%5 == 0 {
        fmt.Print("Buzz")
    }
    
    fmt.Println()
}
