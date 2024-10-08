package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // have type Vertex
	v2 = Vertex{X: 1}  // Y:0 is imlicit
	v3 = Vertex{}      // X:0 & Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
