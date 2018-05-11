package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
	"math"
)

func main() {
	rand2.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Hello")
	//sqrt()
	pointerExample()
}

func rand() int {
	return rand2.Intn(10)
}

func sqrt() {
	fmt.Printf("Square root of a random number is %g", math.Sqrt(float64(rand())))
}

func pointerExample() {
	i := 42
	p := &i
	*p = 21
	fmt.Println(p)
	fmt.Println(&i)
	fmt.Println(*p)
	fmt.Println(&p)
}

func structExample() {
	type Vector struct {
		X int
		Y int
	}

	v := Vector{1, 2}
	p := &v
	p.X = 4 // WHaaaaaat!!, shouldn't it be (*p).X = 4??
	fmt.Println(v)
}
