package main

import "fmt"

func main() {

    var n int
    fmt.Scan(&n)
    
    if n <= 0 {
        fmt.Println("Dimensione non sufficiente")
        return
    }
    
    var triangolo string
    for i := 0; i < n; i++ {
    
        triangolo += "*"
        fmt.Println(triangolo)
    }
}
