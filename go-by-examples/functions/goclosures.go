//uma closure é uma função que está desntro de outra função, assim podendo lembrar as variáveis da função pai, lembrando do "ambiente" em que ela foi criada. Esse tipo de função está ligada mais paradgimas funcionais, mesmo que também possua semelhanças a orientação a objeto, com suas questões de herança

package main

import "fmt"

func intSeq() func() int { //assim nós definimos uma closure
    i := 0
    return func() int { //inline function/função anônima
        i++
        return i
    }
}

func main() {

    nextInt := intSeq()
	nextInt()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt()) //o nextInt irá lembrar quando

    newInts := intSeq()
    fmt.Println(newInts()) //já o newInts é independente do nextInt
}