package main

import "fmt"
import "math"
import "math/rand"

const EPSILON = 1.0e-6

func main() {

    var s int64
    var xC, yC, raggio float64
    
    fmt.Print("Inserisci s: ")
    fmt.Scan(&s)
    fmt.Print("Inserisci i valori dell'ascissa e dell'ordinata del punto C (il centro del cerchio): ")
    fmt.Scan(&xC, &yC)
    fmt.Print("Inserisci il valore del raggio: ")
    fmt.Scan(&raggio)
    
    rand.Seed(s)
    
    var distanzaMax, xMax, yMax float64
    var distanzaMin, xMin, yMin float64 = raggio, xC+raggio, yC+raggio
        //Inizializzati a valori sulla circonferenza
    
    var delta, offsetX, offsetY float64
    delta = raggio * 2
    offsetX = xC - raggio
    offsetY = yC - raggio
    
    for i := 0; i < 1000000; i++ {
        
            //Calcolo random punti con offset e delta (?)
        px := offsetX + (rand.Float64() * delta)
        py := offsetY + (rand.Float64() * delta) 
        
        distanzaP := DistanzaCentro(px, py, xC, yC)
        
        if SulCerchio(distanzaP, raggio) {
            
            fmt.Print("Il punto (", px, ", ", py, ") giace sulla circonferenza del cerchio\n")
            
        } else if DistanzaMaggiore(distanzaP, raggio) {
            //Fuori dal cerchio
            
            if DistanzaMaggiore(distanzaP, distanzaMax) {
            
                distanzaMax = distanzaP
                xMax = px
                yMax = py   
            }
            
        } else {
            //Dentro al cerchio
        
            if DistanzaMinore(distanzaP, distanzaMin) {
                
                distanzaMin = distanzaP
                xMin = px
                yMin = py
            }
        }
    }
    
    fmt.Print("Il punto (", xMax, ", ", yMax, ") è quello all'esterno del cerchio che ha distanza massima dal centro C.\n")
        fmt.Println("Distanza: ", distanzaMax)
        
        fmt.Print("Il punto (", xMin, ", ", yMin, ") è quello all'interno del cerchio che ha distanza minima dal centro C.\n")
        fmt.Println("Distanza: ", distanzaMin)
}

func DistanzaCentro(px, py, xC, yC float64) float64 {
    
    return math.Sqrt(math.Pow((px-xC), 2) + math.Pow((py-yC), 2))
}

func SulCerchio(distanza, raggio float64) bool {
    
        //Controlla se la differenza è irrilevante cioè sono uguali
    return math.Abs(distanza-raggio) < EPSILON
}

/* EQUIVALENTE A DistanzaMaggiore

func EsternoCerchio(distanza, raggio) bool {
        
        //Controlla se la distanza è maggiore del raggio
    return distanza-raggio > EPSILON
} */ 

func DistanzaMaggiore(distanza1, distanza2 float64) bool {
        
        //Controlla se la differenza è significativa (dis1>dis2)
    return distanza1-distanza2 > EPSILON
}

func DistanzaMinore(distanza1, distanza2 float64) bool {

    return distanza2-distanza1 > EPSILON
}
