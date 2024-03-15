package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n //3e20 -> 3^20
    fmt.Println(d) 

    fmt.Println(int64(d)) //arbitrary precision -> constantes fazem cálculos arbitrariamente precisos, logo, infere-se que a resposta será "arredondado"

    fmt.Println(math.Sin(n)) //tira o seno, provavelmente kk
}