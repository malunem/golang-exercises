package main

import "fmt"
import "strings"

func main() {
    
    //var maiuscole string 
    
    for {
    
        var input string 
        fmt.Println("Inserisci una stringa:")
        fmt.Scanln(&input)
        
        if input == "" {
            fmt.Println("ciao")
            break
        }
        
        fmt.Println("Stringa in maiuscolo:", strings.ToUpper(input))
    }
}
