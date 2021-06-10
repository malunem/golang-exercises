package main

import "fmt"
import "math"

func main() {

    var eta1, eta2 int
    
    fmt.Print("Età persona 1 = ")
    fmt.Scan(&eta1)
    
    fmt.Print("Età persona 2 = ")
    fmt.Scan(&eta2)
    
    somma := eta1 + eta2
    media := float64(somma)/2
    mediaInf := math.Floor(media)
    mediaSup := math.Ceil(media)
    somma10 := somma+20
    media10 := float64(somma10)/2
    
    fmt.Print(
        "Somma età = ", somma, "\n" +
        "Media età = ", media, "\n" +
        "Media età arrotondata per difetto all'intero inferiore = ", mediaInf, "\n" +
        "Media età arrotondata per eccesso all'intero superiore = ", mediaSup, "\n" +
        "Somma età a 10 anni = ", somma10, "\n" +
        "Media età a 10 anni = ", media10, "\n")
}
