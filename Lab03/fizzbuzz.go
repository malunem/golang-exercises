package main

import "fmt"

func main() {

    var n int 
    fmt.Scan(&n)
    
    for i := 1; i <= n; i++ {
    
        var fizz bool = (i)%3 == 0
        var buzz bool = (i)%5 == 0
        
        if fizz {
            
            fmt.Print("Fizz ")
        }
        
        if buzz {
        
            fmt.Print("Buzz ")
        }
        
        if !fizz && !buzz {
        
            fmt.Print(i, " ")
        }
    }
    
    fmt.Println()
}
