package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello()
	time.Sleep(1)
	fmt.Println("world")
}
