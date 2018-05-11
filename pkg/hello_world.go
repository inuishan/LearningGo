package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
	"math"
)

func main()  {
	rand2.Seed( time.Now().UTC().UnixNano())
	fmt.Println("Hello")
	sqrt()
}

func rand() int {
	return rand2.Intn(10)
}

func sqrt() {
	fmt.Printf("Square root of a random number is %g", math.Sqrt(float64(rand())))
}