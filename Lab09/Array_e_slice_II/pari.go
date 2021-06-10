package main

import (
        "fmt"
        "os"
        "strconv"
        "math"
)

func main() {
    //legge da riga di comando una sequenza di int os.Args
    var sequenza []int
    
    for _, v := range os.Args[1:] {
        if n, err := strconv.Atoi(v); err == nil {
            sequenza = append(sequenza, n)
        }
    }
    
    //minDispari = numero dispari più basso - func MinimoDispari
    minimoDispari := MinimoDispari(sequenza)
    
    //stampa i valori pari superiori a minDispari
    for _, v := range sequenza {
        if v%2 == 0 && v > minimoDispari {
            fmt.Printf("%d ", v)
        }
    }
    fmt.Println()
}

func MinimoDispari(sl []int) int {
    min := int(math.MaxInt64)
    for _, n := range sl {
        if n%2 != 0 && n < min {
            min = n
        }
    }
    return min
}
