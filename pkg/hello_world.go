package main

import (
	"fmt"
	rand2 "math/rand"
)

func main()  {
	fmt.Println("Hello")
	rand()
}

func rand() {
	fmt.Println("A random number ", rand2.Intn(7))
}