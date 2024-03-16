package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {  //no range de num, ou seja, no tamanho do array de 0 - 2 |||| o _ serve para ignorar o index, algo que não será necessário no momento, mas pode ser para outros casos
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums { //por exemplo aqui, nós iremos utilizar o i para printar o index do valor
        if num == 3 {
            fmt.Println("index:", i, "value:", nums[i])
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for key, value := range kvs {
        fmt.Printf("%s -> %s\n", key, value)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

	word := "go"

    for i, c := range word {
        fmt.Println(i, c)
    }
}