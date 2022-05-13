package test

import "fmt"

func A() {
	fmt.Printf("this is function A()")
}

func B() {
	fmt.Printf("secret string")
}

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) C() {
	fmt.Printf("this is C()")
}

func (v *Vertex) D() {
	fmt.Printf("secret D")
}

/*
func main() {
}
*/
