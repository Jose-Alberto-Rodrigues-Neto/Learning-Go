/*
Crie um programa que simule um menu de opções para um usuário. O menu deve exibir as seguintes opções:
- Calcular área de um círculo: ao escolher pede as informações do círculo e então cálcula
- Calcular área de um retângulo: ao escolher pede as informações do retângulo e então cáclcula
- Sair: termina o programa
*/

package main

import(
	"fmt"
	"math"
)

func main(){
	for{
		fmt.Println("Bem vindo ao program Q2! \nO que você deseja fazer:\n1 - Calcular a área de um círculo\n2 - Calcular a área de um retângulo\n3 - Sair")
		var entrada int
		fmt.Scan(&entrada)
		switch entrada {
		case 1:
			fmt.Println("Entre com o valor do raio:")
			var r float64
			fmt.Scan(&r)
			areaCirculo := math.Pi*r*r
			fmt.Println(areaCirculo)

		case 2:
			fmt.Println("Entre com o valor da base e da altura, separado por espaço (ex: base altura)")
			var a, b float64
			fmt.Scan(&b, &a)
			areaRetangulo :=b*a
			fmt.Println(areaRetangulo)

		case 3:
			fmt.Println("Até mais!")
			return

		default:
			fmt.Println("Insira um valor válido\n")
		}
		
		fmt.Println("Rodar novamente? (s: sim / n: não)")
		var sn string
		fmt.Scan(&sn)
		if sn == "n"{
			break
		}
	}
}