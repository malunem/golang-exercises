package main

import (
        "fmt"
        "strconv"
        "os"
)

func main() {

    n, _ := strconv.Atoi(os.Args[1])
    
    fmt.Println("Fattoriali: ", Fattoriali(n))
}

func Fattoriali(n int) (f []int) {
    
    f = make([]int, n)
    
    for i := 0; i < n; i++ {
    
        f[i] = Fattoriale(i+1)
    }
    
    return f
}

func Fattoriale(n int) (f int) {
    
    if n == 1 {
        return 1
    }
    
    f = n*(Fattoriale(n-1))
    
    return f
}
