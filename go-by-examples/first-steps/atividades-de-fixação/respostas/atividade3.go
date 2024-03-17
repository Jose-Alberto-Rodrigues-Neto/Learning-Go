/*
Crie um programa que armazene uma lista de produtos de um mercado, sendo elas: {banana, maçã, laranja}. O usuário deve digitar o nome de uma fruta, e o programa deve verificar se essa fruta está no estoque, caso ainda esteja pode comprar. 
- Se estiver disponivel compre ou saia; 
- Caso compre, dê a opção de sair, ou olhar outro produto e então tire o elemento do estoque
- Caso contrário, informe que a fruta não foi encontrada e dê a opção de olhar outro produto ou sair.
*/
package main

import "fmt"

func main(){
	
	estoque := []string{"banana", "maçã", "laranja"}

	for{
		fmt.Println("O que você deseja comprar?\n", estoque, "\nDigite 'sair' para sair do programa.")
		var entrada string
		fmt.Scan(&entrada)
		for i, fruta := range estoque{
			if fruta == entrada{
				fmt.Println("Compra bem sucedida!")
				estoque = append(estoque[:i], estoque[i+1:]...)
				break
			}else if entrada == "sair"{
				fmt.Println("Até mais!")
				return
			}else{
				fmt.Println("Esse item não está no estoque!")
			}
		}
		
		fmt.Println("O que você deseja fazer?\n- Para continuar comprando entre com qualquer coisa no terminal.\n- Digite 'sair' para sair.")
		fmt.Scan(&entrada)
		if entrada == "sair"{
			fmt.Println("Até mais!")
			return
		}
	}
}