//funções recursivas chamam elas mesmas

package main

import "fmt"

func fact(n int) int { //nessa função ela só vai retornar o valor quando n == 0
    if n == 0 {
        return 1
    }
    return n * fact(n-1) //aqui a função está chamando ela mesma
}

func main() {
    fmt.Println(fact(7))

    var fib func(n int) int //inicializando a closure

    fib = func(n int) int { //para utilizar recursividade em go é necessário criar a closure utilizando a estrutura da linha 17 e logo após definir o que ela irá fazer, como foi mostrado nessa função anônima
        if n < 2 {
            return n
        }

        return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7))
}