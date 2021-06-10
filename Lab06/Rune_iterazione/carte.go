package main

import "fmt"

func main() {

    var c rune = 127153
    
    for i := 0; i < 10; i++ {
    
        fmt.Printf("%c\n", c)
        c += 1
    }
}
