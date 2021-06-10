package main

import "fmt"

func main() {

    var km float64
    
    fmt.Print("Distanza (Km) = ")
    fmt.Scan(&km)
    
    mi := km * 0.62137
    
    fmt.Println("Distanza (mi) =", mi)
}
