package main

import "fmt"

func main() {

    var ang1, ang2 int
    fmt.Print("Inserire le ampiezze dei due angoli: ")
    fmt.Scan(&ang1, &ang2)
    
    if ang1 + ang2 > 180 {
    
        fmt.Println("I due angoli non appartengono ad un triangolo")

    } else {
    
        fmt.Print("Ampiezza terzo angolo = ", 180 - ang1 - ang2, "° \n")
    }
}
