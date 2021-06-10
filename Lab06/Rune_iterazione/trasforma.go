package main

import "fmt"
import "unicode"

func main() {

    var testo string
    fmt.Scan(&testo)
    
    fmt.Print("Testo maiuscolo: ")
    for _, c := range testo {
    
        fmt.Printf("%c", inMaiuscolo(c))
    }
    
    fmt.Print("\nTesto minuscolo: ")
    for _, c := range testo {
        
        fmt.Printf("%c", inMinuscolo(c))
    }
    fmt.Println()
}

func inMaiuscolo(c rune) rune {

    if c >= 'a' && c <= 'z' {
        c = unicode.ToUpper(c)
    }
    
    return c
}

func inMinuscolo(c rune) rune {

    if c >= 'A' && c <= 'Z' {
        c = unicode.ToLower(c)
    }
    
    return c
}
