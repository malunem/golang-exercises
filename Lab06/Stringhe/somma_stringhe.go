package main

import "fmt"
import "strconv"

func main() {
    
    var somma int
    
    for {
        var num string
        fmt.Println("inserisci un intero:")
        fmt.Scan(&num)
        
        n, err := strconv.Atoi(num)
        
        if err == nil {
                somma += n
        } else {
            fmt.Println("somma:", somma)
            break
        }
    }
}
