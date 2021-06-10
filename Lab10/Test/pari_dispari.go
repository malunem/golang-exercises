package main

import (
        "fmt"
        "strconv"
        "os"
)

func main() {
    n, _ := strconv.Atoi(os.Args[1])
    if Pari(n) {
        fmt.Printf("Il numero %d è pari\n", n)
    } else if Dispari(n) {
        fmt.Printf("Il numero %d è dispari\n", n)
    }
}

func Pari(n int) bool {
    if n%2 == 0 {
        return true
    } else {
        return false
    }
}

func Dispari(n int) bool {
    if Pari(n) {
        return false
    } else {
        return true
    }
}
