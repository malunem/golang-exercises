package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {    
    //Numero perfetto = uguale alla somma dei suoi divisori propri
    //divisori propri = tutti i divisori tranne il numero stesso
   
    //legge da riga di comando 3 int N, DIVISORIMIN, DIVISORIMAX
    N, DIVISORIMIN, DIVISORIMAX := leggiNumeri()
    
    //individua i numeri perfetti e li accoppia con numeri che hanno in comune tot divisori propri, con il numero di divisori compreso tra divisorimin e max
    for a := 0; a <= N; a++ {
        //1. individua tutti i divisori propri
        divisoriA := DivisoriPropri(a)
        //2. controlla se la somma dei divisori è pari al numero
        if èPerfetto(a, divisoriA) {
            for b := 0; b <= N; b++ {
            //3. scorre di nuovo i numeri per trovare num con divisori compatibili
            divisoriB := DivisoriPropri(b)
            divisoriComuni := DivisoriComuni(divisoriA, divisoriB)
            
                if divisoriComuni >= DIVISORIMIN && divisoriComuni <= DIVISORIMAX {
                    fmt.Printf("%d %d\n", a, b)
                }
            }
        }
    }    
    
    //stampa le coppie
}

func leggiNumeri() (N, DIVISORIMIN, DIVISORIMAX int) { 
    val1, err1 := strconv.Atoi(os.Args[1])
        if err1 == nil {
            N = val1
        }
    val2, err2 := strconv.Atoi(os.Args[2])
        if err2 == nil {
            DIVISORIMIN = val2
        }
    val3, err3 := strconv.Atoi(os.Args[3])
        if err3 == nil {
            DIVISORIMAX = val3
        }     
    return N, DIVISORIMIN, DIVISORIMAX
}

func DivisoriPropri(n int) []int {
    //restituisce tutti i divisori propri di n
    var divisori []int
    for i := 1; i < n; i++ {
        if n%i == 0 {
            divisori = append(divisori, i)
        }
    }
    return divisori
}

func èPerfetto(n int, divisori []int) bool {
    var somma int
    for _, v := range divisori {
        somma += v
    }
    
    if somma == n {
        return true    
    } else {
        return false
    }
}

func DivisoriComuni(divisoriA, divisoriB []int) (divisoriComuni int) {
    for _, v1 := range divisoriA {
        for _, v2 := range divisoriB {
            if v1 == v2 {
                divisoriComuni++
            }
        }
    }
    return divisoriComuni
}
