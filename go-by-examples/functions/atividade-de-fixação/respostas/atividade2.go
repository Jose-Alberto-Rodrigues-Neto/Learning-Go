//Crie uma closure em Go que capture uma variável do tipo inteiro e retorne uma função que aceita um inteiro como argumento e verifica se o argumento é menor do que o inteiro capturado.

package main

import "fmt"

func comparaInt(num int) func(int) {
	return func(newNum int) {
		if newNum < num{
			fmt.Println(newNum, "é menor que", num)
		}else if newNum == num {
			fmt.Println(newNum, "é igual a", num)
		}else{
			fmt.Println(newNum, "é maior que", num)
		}
	}
}

func main(){
	fmt.Println("Insira um valor inteiro, não negativo e diferente de zero, que será determinado como o limite!")
	var entrada int
	fmt.Scan(&entrada)
	compara := comparaInt(entrada)
	for{
		fmt.Println("Insira um valor que você deseja comparar: (Para sair digite: 0)")
		fmt.Scan(&entrada)
		if entrada != 0 {
			compara(entrada)
		}else{
			return
		}
	}
}