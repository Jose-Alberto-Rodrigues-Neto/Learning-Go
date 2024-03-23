//Crie uma função em Go chamada media que calcula a média de um conjunto de números, {1.2, 3.3, 5.5, 7.7, 8.8}, fornecidos como argumentos.

package main

import "fmt"

func calMed(nums ...float64) float64{
	var total float64
	for _, num := range nums{
		total = total + num
	}
	return(total/float64(len(nums)))
}

func main(){
	valores := []float64{1.2, 3.3, 5.5, 7.7, 8.8}
	media := calMed(valores...) //é necessário utilizar os 3 pontinhos para ler os valores de dentro do slice
	fmt.Println(media)
}