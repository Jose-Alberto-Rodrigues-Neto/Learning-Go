//Escreva um programa que crie um slice de números inteiros com os valores de 1 a 10. Em seguida, leia os intens do slice e retorne "É par" quando o valor for par.
package main

import "fmt"

func main(){
	slice := []int{1,2,3,4,5,6,7,8,9,10} //cria o slice com os itens já existentes dentro deles
	for _, num := range slice{ //irá percorrer o slice e printar os valores
		if num%2 == 0 { //avisa quando tem um valor par
			fmt.Println(num, "É par")
		}else{
			fmt.Println(num)
		}
	}

	fmt.Println()

	//o equivalente para um array seria
	array := [10]int{1,2,3,4,5,6,7,8,9,10}
	for _, num := range array{
		if num%2 == 0 { //avisa quando tem um valor par
			fmt.Println(num, "É par")
		}else{
			fmt.Println(num)
		}
	}

	fmt.Println()
	
	//sem utilizar range
	for i:=0; i<len(slice); i++{
		if slice[i]%2 == 0 { //avisa quando tem um valor par
			fmt.Println(slice[i], "É par")
		}else{
			fmt.Println(slice[i])
		}
	}
}