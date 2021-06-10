package main

import "fmt"

func main() {

    var n int
    var testo string
    fmt.Scan(&n, &testo)
    
    for i := 0; i < n; i++ {
        
        fmt.Print(testo)
        
        if i < n-1 {
            fmt.Print("-")
        }
    }
    
    fmt.Println()
}
