package main

import "fmt"

func fibonnaci(x int) int { // retorna un int
	if x == 0{
		return 0
	}
	if x == 1{
		return 1
	}
	return fibonnaci(x-1) + fibonnaci(x-2)
}

func main() {
	fmt.Println(fibonnaci(8))
}