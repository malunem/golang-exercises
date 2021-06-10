package main

import (
        "fmt"
        "bufio"
        "os"
        "unicode"
)

func main() {
    
    fmt.Println("Inserisci un testo su più righe (termina con CTRL+D):")
    
    var testo string = LeggiTesto() 
    
    fmt.Printf("Risultato:\n%s\n", TraduciTestoInFarfallino(testo))
}

func LeggiTesto() string {

    var testo string
    fmt.Print()
    
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
    
        if scanner.Text() == "" {
            break
        } else {
            testo += scanner.Text()
        } 
    }
    
    return testo
}

func TraduciTestoInFarfallino(s string) string {
    
    var farfallino string
    
    for _, c := range s {
    
        switch string(c) {
            
            case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
                
                farfallino += TraduciCarattereInFarfallino(c)
            
            default:
                farfallino += string(c)
        }
    }
    
    return farfallino
}


func TraduciCarattereInFarfallino(c rune) string {
    
    //maiuscola := unicode.IsUpper(c)
    
    if unicode.IsUpper(c) {
        return string(c) + string("F") + string(c)
    } else {
        return string(c) + string("f") + string(c)
    }
}
