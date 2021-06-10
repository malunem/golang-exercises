package main

import "fmt"

func main() {

    var testo string
    fmt.Print("Parola in input: ")
    fmt.Scan(&testo)
    
    lunghezza := len(testo)
    
    for i, _ := range testo[0 : (lunghezza+1)/2] {
        
        fmt.Print("Sottostringa ", i, ": ")
        fmt.Printf("%s", testo[i : lunghezza-i])
        fmt.Println()
    }
}
