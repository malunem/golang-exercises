package main

import "fmt"

func main() {

    var n float64
    fmt.Print("Inserisci il valore da arrotondare: ")
    fmt.Scan(&n)
    
    fmt.Println("Valore arrotondato =", float64(int((n*100)+0.5))/100)
}
