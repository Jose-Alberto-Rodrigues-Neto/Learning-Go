package main

import (
    "fmt"
    "maps"
)

func main() {

    m := make(map[string]int) //declara um map, importante para data types associativos

    m["k1"] = 7 //para armazenar um valor deve-se colocar a chave para encontralo e então inicializar o valor
    m["k2"] = 13

    fmt.Println("map:", m) //printa o map

    v1 := m["k1"]
    fmt.Println("v1:", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)

    fmt.Println("len:", len(m))

    delete(m, "k2") //deleta um item
    fmt.Println("map:", m)

    clear(m) //apaga todos os itens do map
    fmt.Println("map:", m)

    _, prs := m["k2"] //aqui nós não iremos utilizar o valor, logo, nós ignoramos com o '_'
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2} //declarando e inicializando um map
    fmt.Println("map:", n)

    n2 := map[string]int{"foo": 1, "bar": 2} 
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }

	maping := make(map[int]string) 
	maping[1] = "primeiro"
	fmt.Println(maping[1])

	//ou maping:= map[int]string{1: primeiro, 2: segundo}
}