//Go supports pointers, allowing you to pass references to values and records within your program.

package main

import "fmt"

func zeroval(ival int) { //essa função não irá mudar o valor do i dentro da main, pois o valor que está dentro da função é uma cópia do valor original
    ival = 0
}

func zeroptr(iptr *int) {  //essa função irá mudar o valor de i dentro da main, pois o parametro da função é um ponteiro, ou seja, a mudança dentro da função vai mudar o valor alocado dentro da memória do parametro selecionado (*int -> define que é um ponteiro)
    *iptr = 0 //utiliza um ponteiro
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i) //aqui nós estamos colocando o endereço da memória de onde o i está armazenado
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)
}