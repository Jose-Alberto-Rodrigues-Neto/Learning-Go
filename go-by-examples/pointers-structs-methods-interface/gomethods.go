package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area() int { //esse é um método que recebe um ponteiro para o struct
    return r.width * r.height
}

func (r rect) perim() int { //esse é um método que recebe um valor para o struct
    return 2*r.width + 2*r.height
}

type dog struct { //criando o objeto cachorro
	name string
	age int
}

func (d *dog) isOld() { //criando um método para o objeto cachorro que fala se o cachorro é jovem ou adulto
	if d.age > 8 {
		fmt.Println("This Dog is old.")
	}else{
		fmt.Println("This Dog is young.")
	}
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())

	d := dog{name: "Mel", age: 14} //criando o objeto cachorro

	d.isOld() //chamando o método
}