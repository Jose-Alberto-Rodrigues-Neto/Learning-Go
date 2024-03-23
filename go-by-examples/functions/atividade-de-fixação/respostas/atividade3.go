//Escreva uma função recursiva em Go chamada potencia que calcula o resultado da potência de um número inteiro base elevado a um expoente inteiro não negativo.

package main

import "fmt"

func potencia(base int, expoente int)  int{
	if expoente == 0 {
		return 1
	}
	return base * potencia(base, expoente-1)
}

func main(){
	var base int
	var expo int
	fmt.Println("Escolha um valor inteiro para ser a base:")
	fmt.Scan(&base)
	fmt.Println("Escolha um valor inteiro para ser o expoente:")
	fmt.Scan(&expo)
	resposta := potencia(base, expo)
	fmt.Println(base, "elevado a", expo, "=", resposta)
}