package main

import "fmt"
import "math"

func CalcolaArea(r float64) float64 {

    return r*r*math.Pi
}

func CalcolaCirconferenza(r float64) float64 {

    return 2*r*math.Pi
}

func main() {

    var raggio float64
    
    fmt.Print("Inserisci il raggio del cerchio: ")
    fmt.Scan(&raggio)
    
    var area float64 = CalcolaArea(raggio)
    var circonferenza float64 = CalcolaCirconferenza(raggio)
    
    fmt.Println("Area del cerchio: ", area)
    fmt.Println("Circonferenza del cerchiio: ", circonferenza)
}
