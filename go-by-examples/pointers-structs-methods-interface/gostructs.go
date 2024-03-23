//continuar dps

package main

import "fmt"

type person struct { //cria uma struct do tipo pessoa
    name string
    age  int
}

func newPerson(name string, age int) *person { //uma função que cria uma nova pessoa, utilizando um ponteiro para person com o objetivo de utilizar os dados fora do scopo da função
    p := person{name: name} //criando o objeto pessoa, apenas com o nome
    p.age = age //atribuindo a idade
    return &p
}

func main() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon", 32))

    s := person{name: "Sean", age: 50} //cria um objeto com nome e idade
    fmt.Println(s.name)

    sp := &s //atribui a sp o endereço do objeto "s"
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
    fmt.Println(s.age)

    dog := struct { //criando uma struct anonimamente
        name   string
        isGood bool
    }{ //passando os dados para o objeto
        "Rex",
        true,
    }

    fmt.Println(dog)

    cat := struct {
        name string
        isGood bool
    }{
        "Moon",
        true,
    }

    fmt.Println(cat)
}