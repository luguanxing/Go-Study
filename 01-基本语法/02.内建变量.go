package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	c := 3 + 4i
	fmt.Printf("Abs(3+4i)=%f\n", cmplx.Abs(c))
	fmt.Println("e^i(Pi)+1=", cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("e^i(Pi)+1=%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	euler();
	triangle();
}
