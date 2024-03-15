package main

import "fmt"
	
func main() {

    var a [5]int //declarando um array de int
    fmt.Println("emp:", a) //lendo o array inteiro

    a[4] = 100 //iniciando valor para um index dentro do array
    fmt.Println("set:", a)
    fmt.Println("get:", a[4]) //lendo o valor do quarto item do array

    fmt.Println("len:", len(a)) //lendo o tamanho do array

    b := [5]int{1, 2, 3, 4, 5} //declarando e iniciando um array de cinco intens 
    fmt.Println("dcl:", b)

	const colunas, linhas = 3, 10 //declarando os valores de linhas e colunas
		
    var twoD [colunas][linhas]int //declarando uma matriz
    for i := 0; i < colunas; i++ { //O(n^2) -> for dentro de for
        for j := 0; j < linhas; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD) //lendo a matriz que foi criada automaticamente por duas interações for
}