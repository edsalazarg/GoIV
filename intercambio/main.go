package main

import "fmt"
  
func variadic(a,b *int){
	temporal := *a
	*a = *b
	*b = temporal
}

func main() {
	var a,b int
	a = 5
	b = 10
	variadic(&a,&b)
	fmt.Println(a)
	fmt.Println(b)
}