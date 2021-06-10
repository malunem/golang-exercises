package main

import "fmt"

func main() {
    
    var (
    n1, n2 int
    sum int
    )
    
    fmt.Scan(&n1, &n2)
    
    for i := n1; i < n2; i++ {
        
        if i %2 == 0 {
            sum += i
        }
    }

    fmt.Println("Somma =", sum)
}
