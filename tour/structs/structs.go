package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)

	// pointers to structs
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// struct literals
	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	p = &Vertex{1, 2}  // has type *Vertex
	fmt.Println(v1, p, v2, v3)
}
