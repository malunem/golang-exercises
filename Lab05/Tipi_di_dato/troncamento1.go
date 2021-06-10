package main

import "fmt"

func main() {

    var n float64
    fmt.Print("Inserisci il valore da troncare: ")
    fmt.Scan(&n)
    
    fmt.Println("Valore troncato =", float64(int(n*100))/100)
}
