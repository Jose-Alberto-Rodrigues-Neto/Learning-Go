package main

import "fmt"

func main() {

    var a = "initial" //string
    fmt.Println(a)

    var b, c int = 1, 2 //int
    fmt.Println(b, c)

    var d = true //bool
    fmt.Println(d)

    var e int //quando um valor não é inicializado, ele recebe o valor "zero"
    fmt.Println(e)

    f := "apple" //é uma versão curta para declarar e inicializar uma variável, sendo possível apenas dentro de funções (como a função "main")
    fmt.Println(f)
}