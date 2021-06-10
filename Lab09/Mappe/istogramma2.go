package main

import (
        "fmt"
        "bufio"
        "os"
        "unicode"
)

func main() {
    //legge os.Stdin fino EOF
    testo := LeggiTesto()
    //crea mappa con occorrenze lettere case sensitive
    occorrenze := CreaMappa(testo)
    //stampa con func apposita
    StampaIstogramma(occorrenze)
}
func LeggiTesto() (testo string) {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        testo += scanner.Text()
    }
    return testo
}

func CreaMappa(testo string) map[rune]int {
    occorrenze := map[rune]int{}
    for _, v := range testo {
        if unicode.IsLetter(v) {
            occorrenze[v]++
        }
    }
    return occorrenze
}

func StampaIstogramma(occorrenze map[rune]int) {
    //ordina le rune in ordine crescente
    var lettere []rune
    for i, _ := range occorrenze {
        lettere = append(lettere, i)
    }
    SortRunes(lettere)
    //stampa * ripetuto in base a occorrenze in ordine crescente di unicode
    for _, v := range lettere {
        fmt.Printf("%c: ", v)
        for j := 0; j < occorrenze[v]; j++ {
            fmt.Printf("*")
        }
        fmt.Println()
    }
    
}

func SortRunes(a []rune) {
    lunghezza := len(a)
	for i := 0; i < lunghezza - 1; i++ {
		indiceMin := i
		for j := i + 1; j < lunghezza - 1; j++ {
			if a[indiceMin] > a[j] {
				indiceMin = j
			}
		}
		a[i], a[indiceMin] = a[indiceMin], a[i]
	}
}
