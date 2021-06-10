package main

import "fmt"

func main() {

    var parola string
    fmt.Scan(&parola)
    
    if èPalindroma(parola) {
        fmt.Println("Palindroma")
    } else {
        fmt.Println("Non palindroma")
    }
}

func èPalindroma(parola string) bool {
    
    lunghezza := len(parola)
    
    for i := range parola {
        
        if parola[i] != parola[lunghezza-1-i] {
            return false
        }
    }
    
    return true
}
