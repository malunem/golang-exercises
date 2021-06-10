package main

import (
        "fmt"
        "bufio"
        "os"
        "unicode"
)

func main() {
    
    var testo string 
    testo = LeggiTesto()
    testo = Spazia(testo)
    
    fmt.Printf("Risultato: %s\n", testo)
}

func LeggiTesto() string {
    //usare bufio.Scanner con os.Stdin fino EOF
    
    var input string 
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        
        input += scanner.Text()
    }
    
    return input
}

func Spazia(s string) string {
    //inserisce spazio tra due caratteri non spazi (unicode.IsSpace)
    
    var spazi string
    var lastChar bool = true //true se è uno spazio 
    
    for _, c := range s {
        
        switch {
            
            case !unicode.IsSpace(c) && !lastChar:
                spazi += string(" ")
                spazi += string(c)
                lastChar = false
                
            case !unicode.IsSpace(c) && lastChar:
                spazi += string(c)
                lastChar = false
            
            case unicode.IsSpace(c):
                spazi += string(c)
                lastChar = true
        }
    }
    
    return spazi
}
