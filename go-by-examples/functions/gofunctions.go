package main

import "fmt"

func sum(a int, b int) int { //func nomeFunc(param type, param type) type{ body  return (something) }

    return a + b
}

func sumSumDiv(nums ...int) (int, int) { //função que retorna mais de um valor e pode receber mais de um valor (nums ...int) -> é arbitrário
    total := 0
    for _, num := range nums{
        total = total + num
    }
    return total, (total)/2
}

func unit(name string) { //func que não retorna nada
	fmt.Println(name, "Unit")
}

func main() {

    res := sum(1, 2)
    fmt.Println("1+2 =", res)
    res1, resDiv := sumSumDiv(1, 2, 3) //para retornar mais de um valor é essencial que tenha mais de uma variavel declarada para ser iniciada, ou então utilizar '_' em uma delas para ignorar o valor
    fmt.Println("1+2+3 =", res1, "(1+2+3)/2 =", resDiv)
	unit("Func")
}