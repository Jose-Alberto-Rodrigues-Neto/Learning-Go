package main

import "fmt"

func main() {

	// Single condition
	fmt.Println() //br line
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after for loop
	fmt.Println() //br line
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// "for" without a condition will loop repeatedly
	// until you "break" out of the loop
	// or "return" from the enclosing function
	fmt.Println() //br line
	for {
		fmt.Println("loop")
		break
	}

	// You can also "continue" to the next iteration of the loop
	fmt.Println() //br line
	for n := 0; n <= 5; n++ { //não lê os números pares
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}