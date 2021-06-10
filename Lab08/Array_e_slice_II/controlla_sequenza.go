package main

import (
        "fmt"
        "unicode"
)

func main() {

    var input string
    var numeri string
    fmt.Scan(&input)
    
    for _, v := range input {
    
        if unicode.IsDigit(v) {
        
            numeri += string(v)
        }
    }
    
    for i := range numeri[:len(numeri)-2] {
        
        if numeri[i] < numeri[i+1] {
            fmt.Printf("Sequenza nascosta non ordinata.\n")
            return
        }
    }
    
    fmt.Printf("Sequenza nascosta ordinata.\n")
}
