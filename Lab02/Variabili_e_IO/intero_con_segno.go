package main

import "fmt"

func main() {
    var n int
    
    fmt.Print("Inserisci numero: ")
    fmt.Scan(&n)
    
/*    if n > 0 {
        fmt.Print("+", n, "\n")
    } else {
        fmt.Println(n)
    }
} */

    if n > 0 {
        fmt.Print("+")
    }
    
    fmt.Print(n, "\n")
}
