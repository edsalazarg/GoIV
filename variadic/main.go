package main

import "fmt"
  
func variadic(args ...int) int{
	maximo := 0
	for _, v := range args { 
		if v > maximo{
			maximo = v
		}
	}
	return maximo
}

func main() {
	maximo := variadic(2,3,15,4,7,8,9,15,10,13,11,10,12,4,5,6,9)

	fmt.Println(maximo)
	
}