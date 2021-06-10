package main

import (
        "fmt"
        "os"
        "strconv"
        "math/rand"
        "time"
)

func main() {
    
    n := LeggiNumero()
    
    if n <= 0 || n >= 10 {
        fmt.Println("Inserisci un numero compreso tra 0 e 10")
        return
    }
    
    mazzo := GeneraMazzo()
    
    rand.Seed(int64(time.Now().Nanosecond()))
    
    EstraiCarte(mazzo, n)
}

func LeggiNumero() int {

    n, _ := strconv.Atoi(os.Args[1])
    return n
}

func GeneraMazzo() string {

    var mazzo string
    
    for i := 0; i < 10; i++ {
        
        mazzo += string('\U0001F0B1' + int32(i))
    } 
    
    return mazzo
}

func EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string) {
    
    mazzoRune := []rune(mazzo)
    numeroCarta := rand.Intn(len(mazzoRune)) //'\U0001F0B1'
    
    for i, v := range mazzoRune {
        
        if i != numeroCarta {
            
            mazzoResiduo += string(v)
            
        } else if i == numeroCarta {
            
            cartaEstratta = v
        }
    }
    
    return
}

func EstraiCarte(mazzo string, n int) {
    
    var cartaEstratta rune
    var mazzoResiduo string = mazzo
    
    for i := 0; i < n; i++ {
    
        cartaEstratta, mazzoResiduo = EstraiCarta(mazzoResiduo)
        
        fmt.Printf("Estratta la carta %c - Carte rimaste nel mazzo: %v\n\n", cartaEstratta, mazzoResiduo)
    }
}
