package main

import "fmt"
import "math"

const EPSILON = 1.0e-9

func main() {

    var xA, yA, xB, yB, xC, yC float64
    fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto A: ")
    fmt.Scan(&xA, &yA)
    fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto B: ")
    fmt.Scan(&xB, &yB)
    fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto A: ")
    fmt.Scan(&xC, &yC)
    
    fmt.Println()
    
    var AB, BC, CA float64
    
    AB = Distanza(xA, yA, xB, yB)
    BC = Distanza(xB, yB, xC, yC)
    CA = Distanza(xC, yC, xA, yA)
    
    var ABeBC, BCeCA, CAeAB bool
    
    ABeBC = LatiUguali(AB, BC)
    BCeCA = LatiUguali(BC, CA)
    CAeAB = LatiUguali(CA, AB)
    
    if ABeBC && BCeCA && CAeAB {
        
        //equilatero
        
        fmt.Println("Il triangolo ABC è equilatero.")
        fmt.Println("Lunghezza del lato:", AB)
        
    } else if ABeBC || BCeCA || CAeAB {
    
        //isoscele
        
        fmt.Println("Il triangolo ABC è isoscele.")
        fmt.Print("I lati di lunghezza uguale sono ")
        
        if ABeBC {
            
            fmt.Print("AB e BC.\n") 
            fmt.Println("Lunghezza dei lati AB e BC:", AB)
            fmt.Println("Lunghezza del lato CA:", CA)
        
        } else if BCeCA {
        
            fmt.Print("BC e CA.\n")
            fmt.Println("Lunghezza dei lati BC e CA:", BC)
            fmt.Println("Lunghezza del lato AB:", AB)
            
        } else {
        
            fmt.Print("CA e AB.\n")
            fmt.Println("Lunghezza dei lati CA e AB:", CA)
            fmt.Println("Lunghezza del lato BC:", BC)
        
        }
        
        
    } else {
    
        //scaleno
        
        fmt.Println("Il triangolo ABC è scaleno.")
        fmt.Println("Lunghezza del lato AB:", AB)
        fmt.Println("Lunghezza del lato BC:", BC)
        fmt.Println("Lunghezza del lato CA:", CA)
        
    }
}

func Distanza(x1, y1, x2, y2 float64) float64 {

    distanza := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
    
    return distanza
}

func LatiUguali(lato1, lato2 float64) bool {

    return math.Abs(lato1-lato2) <= EPSILON
}
