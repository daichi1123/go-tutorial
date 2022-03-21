package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (v *Vertex) String() string {
	return fmt.Sprintf("x is %d! Y is %d!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
}
