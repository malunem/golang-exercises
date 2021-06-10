package main

import "fmt"

func main() {

    //scrivere i numeri da 1 a 100 con due cicli: decine e unità
    
    var dec, uni int
    
    for dec = 0; dec < 10; dec++ {
    
        for uni = 0; uni < 10; uni++ {
            
            fmt.Println(dec*10 + uni)
        }
    }
}
