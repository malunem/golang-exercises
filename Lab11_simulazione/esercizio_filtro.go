package main

import (
        "fmt"
        //"bufio"
        "os"
        "strconv"
        //"unicode"
)

func main() {
    
    //Legge da riga di comando ch n
    input := os.Args[1:]
    
    var output string
    
    for i := 0; i < len(input); i+=2 {
        
        if n, err := strconv.Atoi(input[i+1]); err == nil {
            
            //Ripete n volte la stringa ch concatenandola
            for j := 0; j < n; j++ {
                
                output += input[i]
            }
        }
    }
    
    //Stampa la stringa ottenuta
    fmt.Printf("%s\n", output)
}
