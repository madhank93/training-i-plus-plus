package main

import "fmt"

func add(a int, b int) int{
	return a+b
}

func main(){
	fmt.Println("Sum is: ", add(20,30))
}