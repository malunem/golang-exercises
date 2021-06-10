package main

import "fmt"

func main() {
    
    var n int
    fmt.Scan(&n)
    
    if n == 0 {
        fmt.Println("Dimensione non sufficiente")
        return
    }
    
    diagonale := 2*n+1
    
    var stringaSpazi, stringaAsterischi string
    
    for i := 0; i < diagonale; i++ {
    
        if i < n {
            stringaSpazi += " "
        }
        
        stringaAsterischi += "*"
    }
    
    for i := 0; i < n+1; i++ {
    
        fmt.Print(stringaSpazi[i:], stringaAsterischi[:i*2+1], "\n")   
    } 
    
    for i := n; i > 0; i-- {
    
        fmt.Print(stringaSpazi[i-1:], stringaAsterischi[:i*2-1], "\n")
    }
}
