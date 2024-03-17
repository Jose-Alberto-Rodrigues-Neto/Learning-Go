package main

import (
    "fmt"
    "slices" //serve para as funções adjacentes relacionadas a slices: como slices.Equal (https://go.dev/blog/slices-intro)
)

func main() {

    var s []string //são declaradas sem valores, tendo um tamanho de zero e valores "nil" (nil é referente ao valor default de uma variavel não inicializada)
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    s = make([]string, 3) //inicializa uma slice com tamanho definido utilizando o make([]string, tamSlice)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))

    s = append(s, "d")
    s = append(s, "e", "f") //pode adicionar mais de um elemento no slice de uma só vez utilizando append
    fmt.Println("apd:", s)

    c := make([]string, len(s)) //cria uma slice de tamanho igual ao de "s"
    copy(c, s) //popula a slice c com os dados da slice s
    fmt.Println("cpy:", c)

    l := s[2:5] //seleciona os dados que estão a cima do indice head até o indice da tail, escluindo os mesmos ->  ]head -> tail[
    fmt.Println("sl1:", l)

    l = s[:5] // (h:t[; onde t=5
    fmt.Println("sl2:", l)

    l = s[2:] // ]h:t); onde h =2
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"} //inicializado e declarado instantaneamente
    fmt.Println("dcl:", t)

    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) { //comparação de slices
        fmt.Println("t == t2")
    }

    twoD := make([][]int, 3) //cria uma matriz
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen) //cria o tamanho do slice "dentro do slice"
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}