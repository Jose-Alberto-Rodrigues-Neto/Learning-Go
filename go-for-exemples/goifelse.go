package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 é par") //par em inglês é even
    } else {
        fmt.Println("7 é impar") //impar em inglês é odd
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

    if num := 9; num < 0 { //achei legal isso, o número pode ser declarado como uma váriavel dentro do if e dps reutilizado dentro do escopo do if!
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	//obs: não existe if ternário em Go!
}