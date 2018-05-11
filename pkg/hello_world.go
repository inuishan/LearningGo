package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

func main()  {
	rand2.Seed( time.Now().UTC().UnixNano())
	fmt.Println("Hello")
	rand()
}

func rand() {
	fmt.Println("A random number ", rand2.Intn(10))
}