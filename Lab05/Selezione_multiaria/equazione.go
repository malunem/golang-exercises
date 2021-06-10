package main

import "fmt"

func main() {

    var a, b int
    fmt.Scan(&a, &b)
    
    var x float64
    x = -float64(b)/float64(a)
    
    fmt.Println("La radice è ", x)
}
