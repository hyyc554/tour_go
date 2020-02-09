package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var abser Abser
	f := MyFloat(-100.25)
	abser = f
	fmt.Println(abser.Abs())
	v := Vertex{10, 20}
	abser = &v
	fmt.Println(abser.Abs())
	a := Vertex{3, 4}
	fmt.Println(a.Abs())
}
